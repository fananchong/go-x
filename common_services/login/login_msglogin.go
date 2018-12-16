package main

import (
	"crypto/md5"
	"fmt"
	"net/http"

	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/fananchong/go-x/base"
	service "github.com/fananchong/go-x/common_services"
	"github.com/fananchong/go-x/common_services/db"
	"github.com/fananchong/go-x/common_services/proto"
	"github.com/fananchong/go-x/internal/common/k8s"
	proto1 "github.com/golang/protobuf/proto"
	uuid "github.com/satori/go.uuid"
)

var curIndex = 0

func (this *Login) MsgLogin(w http.ResponseWriter, req *http.Request, data string, sign string) {
	msg := &proto.MsgLogin{}
	if this.decodeMsg(data, msg) == nil {
		w.Write(getErrRepString(proto.EnumLogin_ErrParams))
		return
	}
	base.XLOG.Debugln("account =", msg.GetAccount())
	base.XLOG.Debugln("password =", msg.GetPassword())
	base.XLOG.Debugln("mode =", msg.GetMode())

	var accountId uint64
	var password string
	var err error
	if int(msg.GetMode()) < int(proto.LoginMode_CUSTOM_BEGIN) {
		// common层支持的登录模式
		switch msg.GetMode() {
		case proto.LoginMode_Default:
			accountId, password, err = this.loginByDefault(msg)
			if err != nil {
				w.Write(getErrRepString(proto.EnumLogin_ErrDB))
				return
			}
			// 密码可以为空，因此不需要做否为空的判断
		default:
			base.XLOG.Debugln("unknow mode, mode =", msg.GetMode())
			w.Write(getErrRepString(proto.EnumLogin_ErrMode))
			return
		}
	} else {
		// 应用层的登录模式
		passwd, err := this.GetPassword(msg.GetAccount(), msg.GetMode(), msg.GetUserdata())
		if err != proto.EnumLogin_NoErr {
			w.Write(getErrRepString(err))
			return
		}
		if passwd == "" {
			w.Write(getErrRepString(proto.EnumLogin_ErrPlatformSide))
			return
		}
		password = passwd
	}

	if password != "" && !this.checkPassword(msg.GetPassword(), password, checkSalt(msg.GetMode())) {
		w.Write(getErrRepString(proto.EnumLogin_ErrPassword))
		return
	}

	if accountId == 0 {

		// TODO: 帐号名合法性检查

		uid, err := this.suid.New(db.SUID_TYPE_ACCOUNT)
		if err != nil {
			w.Write(getErrRepString(proto.EnumLogin_ErrDB))
			return
		}
		account := db.NewAccount(this.dbAccountName, msg.GetAccount())
		account.SetUid(uid)
		account.SetPswd(msg.GetPassword())
		err = account.Save()
		if err != nil {
			w.Write(getErrRepString(proto.EnumLogin_ErrDB))
			return
		}
		accountId = uid
	}

	// 获取一个Gateway
	endpoints, err := k8s.GetEndpoints(k8s.GetNamespace(int(base.Gateway)), k8s.GetServiceName(int(base.Gateway)))
	if err != nil || len(endpoints) == 0 {
		w.Write(getErrRepString(proto.EnumLogin_ErrGateway))
		return
	}
	endpointIndex := curIndex % len(endpoints)
	curIndex++

	// 生成Token、保存Token
	temptkn := ""
	temptkn = uuid.NewV4().String()

	token := db.NewToken(this.dbTokenName, msg.GetAccount())
	token.Expire(60 * 30) // 30分钟
	token.SetUid(accountId)
	token.SetToken(temptkn)
	err = token.Save()
	if err != nil {
		w.Write(getErrRepString(proto.EnumLogin_ErrDB))
		return
	}

	// 登录成功
	endpoint := endpoints[endpointIndex]
	ip := endpoint.IP
	iplist := service.GetIpList()
	if v, ok := (*iplist)[ip]; ok {
		ip = v
	}
	addr := fmt.Sprintf("%s:%d", ip, endpoint.Ports[""])
	base.XLOG.Debugln("accountId =", accountId)
	base.XLOG.Debugln("gateway address =", addr)
	rep := &proto.MsgLoginResult{}
	rep.Err = proto.EnumLogin_NoErr
	rep.Token = temptkn
	rep.Address = addr
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

func getErrRepString(err proto.EnumLogin_Error) []byte {
	base.XLOG.Debugln("err =", err)
	rep := &proto.MsgLoginResult{}
	rep.Err = err
	data, _ := proto1.Marshal(rep)
	return data
}

func (this *Login) checkPassword(msgPassword, dbPassword string, isSalt bool) bool {
	if isSalt {
		s1 := []byte(externArgs.Sign1 + dbPassword + externArgs.Sign2 + base.XARGS.Common.Version)
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
