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
	cmds    map[proto.MsgTypeCmd]LoginMsgHandlerType
	Derived ILogin

	// sign
	sign1   string
	sign2   string
	sign3   string
	version string
}

func (this *Login) Start(addr string) {
	if this.cmds == nil {
		this.cmds = make(map[proto.MsgTypeCmd]LoginMsgHandlerType)
		this.cmds[proto.MsgTypeCmd_Login] = this.MsgLogin
	}
	pb.SetLogger(common.GetLogger())
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

func (this *Login) SetSign1(sign string) {
	this.sign1 = sign
}

func (this *Login) SetSign2(sign string) {
	this.sign2 = sign
}

func (this *Login) SetSign3(sign string) {
	this.sign3 = sign
}

func (this *Login) SetVersion(ver string) {
	this.version = ver
}
