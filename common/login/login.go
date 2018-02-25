package login

import (
	"net/http"

	pb "github.com/fananchong/go-proto-helper"
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/proto"
)

type ILogin interface {
	GetPassword(string, proto.LoginMode, []byte) (string, proto.LoginError)
}

type LoginMsgHandlerType func(http.ResponseWriter, *http.Request, string, string)

type Login struct {
	common.WebService
	db      *common.RedisObj
	cmds    map[proto.MsgTypeCmd]LoginMsgHandlerType
	Derived ILogin
}

func (this *Login) Start() bool {
	if this.cmds == nil {
		this.cmds = make(map[proto.MsgTypeCmd]LoginMsgHandlerType)
		this.cmds[proto.MsgTypeCmd_Login] = this.MsgLogin
	}
	this.db = common.NewRedisObj(common.GetArgs().DbAccount.Name, common.GetArgs().DbAccount.Addrs)
	if this.db == nil {
		return false
	}
	pb.SetLogger(common.GetLogger())
	this.HandleFunc("/msg", this.request)
	this.ListenAndServe(common.GetArgs().Login.Listen)
	return true
}

func (this *Login) Register(cmd proto.MsgTypeCmd, f LoginMsgHandlerType) {
	if _, ok := this.cmds[cmd]; !ok {
		this.cmds[cmd] = f
	} else {
		panic("Register fail.")
	}
}
