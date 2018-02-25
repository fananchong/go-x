package login

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/fananchong/go-x/common"
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

	if int(msg.GetMode()) < int(proto.LoginMode_CUSTOM_BEGIN) {
		// common层支持的登录模式
		switch msg.GetMode() {
		case proto.LoginMode_Default:
			accountId, password = this.loginByDefault(msg)
			// 密码可以为空，因此不需要做否为空的判断
		default:
			common.GetLogger().Debugln("unknow mode, mode =", msg.GetMode())
			w.Write(getErrRepString(proto.LoginError_ErrMode))
			return
		}
	} else {
		// 应用层的登录模式
		passwd, err := this.Derived.GetPassword(msg.GetAccount(), msg.GetMode(), msg.GetUserdata())
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

		// 访问本地数据库，获取帐号ID
		// 如果数据不存在，创建帐号数据、角色数据（注意数据库需要原子操作）
		// 获取帐号ID

	}

	// 生成Token、保存Token

	// 登录成功
	common.GetLogger().Debugln("accountId =", accountId)
	w.Write(getErrRepString(proto.LoginError_NoErr))
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
		s1 := []byte(common.GetArgs().Login.Sign1 + dbPassword + common.GetArgs().Login.Sign2 + common.GetArgs().Common.Version)
		s2 := md5.Sum(s1)
		s3 := fmt.Sprintf("%x", s2)
		return msgPassword == s3
	} else {
		return msgPassword == dbPassword
	}
}

func (this *Login) loginByDefault(msg *proto.MsgLogin) (accountId uint64, password string) {
	// mode: default， 访问本地数据库
	// 获取帐号ID、密码
	return
}
