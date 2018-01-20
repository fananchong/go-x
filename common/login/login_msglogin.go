package login

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/proto"
)

func (this *Login) MsgLogin(w http.ResponseWriter, req *http.Request, data string, sign string) {
	msg := this.decodeMsg(data, &proto.MsgLogin{}).(*proto.MsgLogin)
	common.GetLogger().Debugln("account =", msg.GetAccount())
	common.GetLogger().Debugln("password =", msg.GetPassword())
	common.GetLogger().Debugln("mode =", msg.GetMode())

	var accountId uint64
	var password string

	if int(msg.GetMode()) < int(proto.LoginMode_Unknow) {
		switch msg.GetMode() {
		case proto.LoginMode_Default:
			accountId, password = this.loginByDefault(msg)
		default:
			common.GetLogger().Debugln("unknow mode, mode =", msg.GetMode())
			return
		}
	} else {
		passwd, err := this.Derived.GetPassword(msg.GetAccount(), msg.GetMode(), msg.GetUserdata())
		if err != proto.LoginError_NoErr {
			// TODO:
			return
		}
		if passwd == "" {
			// TODO:
			return
		}
		password = passwd
	}

	if password != "" && !this.checkPassword(msg.GetPassword(), password, msg.GetIsSalt()) {
		// TODO:
		return
	}

	if accountId == 0 {

		if msg.GetMode() == proto.LoginMode_Default && msg.GetIsSalt() {
			// TODO:
			return
		}

		// 访问本地数据库，获取帐号ID
		// 如果数据不存在，创建帐号数据、角色数据（注意数据库需要原子操作）
		// 获取帐号ID
	}

	// 生成Token、设置Cookie、保存Token

	// 登录成功

	common.GetLogger().Debugln("accountId =", accountId)
}

func (this *Login) checkPassword(msgPassword, dbPassword string, isSalt bool) bool {
	if isSalt {
		s1 := []byte(this.sign1 + dbPassword + this.sign2 + this.version)
		s2 := md5.Sum(s1)
		s3 := fmt.Sprintf("%x", s2)
		return msgPassword == s3
	} else {
		// 不开放原始密码比对，以提高安全性
		return false
	}
}

func (this *Login) loginByDefault(msg *proto.MsgLogin) (accountId uint64, password string) {
	// mode: default， 访问本地数据库
	// 获取帐号ID、密码
	return
}
