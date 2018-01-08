package login

import (
	"net/http"

	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/proto"
)

type ILogin interface {
	OnVerifyAccount()
}

type LoginMsgHandlerType func(http.ResponseWriter, *http.Request, string)

type Login struct {
	common.WebService
	cmds    map[proto.MsgTypeCmd]LoginMsgHandlerType
	Derived ILogin
}

func (this *Login) Start(addr string) {
	if this.cmds == nil {
		this.cmds = make(map[proto.MsgTypeCmd]LoginMsgHandlerType)
		this.cmds[proto.MsgTypeCmd_Login] = this.MsgLogin
	}
	this.HandleFunc("/msg", this.request)
	this.ListenAndServe(addr)
}

func (this *Login) Register(cmd proto.MsgTypeCmd, f LoginMsgHandlerType) {
	if _, ok := this.cmds[cmd]; !ok {
		this.cmds[cmd] = f
	} else {
		panic("Register fail.")
	}
}
