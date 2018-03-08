package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"github.com/satori/go.uuid"

	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/db"
	"github.com/fananchong/go-x/common/discovery"
	"github.com/fananchong/go-x/common/proto"
	proto1 "github.com/golang/protobuf/proto"
)

func (this *Login) MsgLogin(w http.ResponseWriter, req *http.Request, data string, sign string) {
	msg := this.decodeMsg(data, &proto.MsgLogin{}).(*proto.MsgLogin)
	common.GetLogger().Debugln("account =", msg.GetAccount())
	common.GetLogger().Debugln("password =", msg.GetPassword())
	common.GetLogger().Debugln("mode =", msg.GetMode())

	var accountId uint64
	var password string
	var err error
	if int(msg.GetMode()) < int(proto.LoginMode_CUSTOM_BEGIN) {
		// common层支持的登录模式
		switch msg.GetMode() {
		case proto.LoginMode_Default:
			accountId, password, err = this.loginByDefault(msg)
			if err != nil {
				w.Write(getErrRepString(proto.LoginError_ErrDB))
				return
			}
			// 密码可以为空，因此不需要做否为空的判断
		default:
			common.GetLogger().Debugln("unknow mode, mode =", msg.GetMode())
			w.Write(getErrRepString(proto.LoginError_ErrMode))
			return
		}
	} else {
		// 应用层的登录模式
		passwd, err := this.GetPassword(msg.GetAccount(), msg.GetMode(), msg.GetUserdata())
		if err != proto.LoginError_NoErr {
			w.Write(getErrRepString(err))
			return
		}
		if passwd == "" {
			w.Write(getErrRepString(proto.LoginError_ErrPlatformSide))
			return
		}
		password = passwd
	}

	if password != "" && !this.checkPassword(msg.GetPassword(), password, checkSalt(msg.GetMode())) {
		w.Write(getErrRepString(proto.LoginError_ErrPassword))
		return
	}

	if accountId == 0 {
		uid, err := this.suid.New(db.SUID_TYPE_ACCOUNT)
		if err != nil {
			w.Write(getErrRepString(proto.LoginError_ErrDB))
			return
		}
		account := db.NewAccount(this.dbAccountName, msg.GetAccount())
		account.SetUid(uid)
		account.SetPswd(msg.GetPassword())
		err = account.Save()
		if err != nil {
			w.Write(getErrRepString(proto.LoginError_ErrDB))
			return
		}
		accountId = uid
	}

	// 获取一个Gateway
	gw, _ := discovery.GetNode().Servers.GetOne(int(common.Gateway))
	if gw == nil {
		w.Write(getErrRepString(proto.LoginError_ErrGateway))
		return
	}

	// 生成Token、保存Token
	temptkn := ""
	uid, err := uuid.NewV4()
	if err == nil {
		temptkn = uid.String()
	} else {
		temptkn = fmt.Sprintf("%d%d", time.Now().UnixNano()*1234, accountId*2345)
	}

	token := db.NewToken(this.dbTokenName, msg.GetAccount())
	token.Expire(60 * 30) // 30分钟
	token.SetUid(accountId)
	token.SetToken(temptkn)
	err = token.Save()
	if err != nil {
		w.Write(getErrRepString(proto.LoginError_ErrDB))
		return
	}

	// 登录成功
	common.GetLogger().Debugln("accountId =", accountId)
	common.GetLogger().Debugln("gateway =", gw.GetExternalIp())
	rep := &proto.MsgLoginResult{}
	rep.Err = proto.LoginError_NoErr
	rep.Token = temptkn
	rep.Address = gw.GetExternalIp()
	succmsg, _ := proto1.Marshal(rep)
	w.Write(succmsg)
}

// 帐号密码通过第3方平台获取的，密码都必须salt。
func checkSalt(mode proto.LoginMode) bool {
	if mode == proto.LoginMode_Default {
		return false
	}
	return true
}

func getErrRepString(err proto.LoginError) []byte {
	common.GetLogger().Debugln("err =", err)
	rep := &proto.MsgLoginResult{}
	rep.Err = err
	data, _ := proto1.Marshal(rep)
	return data
}

func (this *Login) checkPassword(msgPassword, dbPassword string, isSalt bool) bool {
	if isSalt {
		s1 := []byte(xargs.Login.Sign1 + dbPassword + xargs.Login.Sign2 + xargs.Common.Version)
		s2 := md5.Sum(s1)
		s3 := fmt.Sprintf("%x", s2)
		return msgPassword == s3
	} else {
		return msgPassword == dbPassword
	}
}

func (this *Login) loginByDefault(msg *proto.MsgLogin) (uint64, string, error) {
	account := db.NewAccount(this.dbAccountName, msg.GetAccount())
	err := account.Load()
	if err != nil {
		if err == go_redis_orm.ERR_ISNOT_EXIST_KEY {
			return 0, "", nil
		} else {
			return 0, "", err
		}
	}
	return account.GetUid(), account.GetPswd(), nil
}
