package main

import (
	"github.com/fananchong/go-x/common/login"
	"github.com/fananchong/go-x/common/proto"
)

var (
	xlogin *ServiceLogin = nil
)

type ServiceLogin struct {
	login.Login
}

func NewServiceLogin() *ServiceLogin {
	this := &ServiceLogin{}
	this.Derived = this
	this.Login.SetSign1("5UY6$f$h")
	this.Login.SetSign2("3wokZB%q")
	this.Login.SetSign3("%2Fi9TRf")
	this.Login.SetVersion(xargs.ArgsBase.Version)
	return this
}

// 第3方平台，获取帐号密码
func (this *ServiceLogin) GetPassword(account string, mode proto.LoginMode, userdata []byte) (string, proto.LoginError) {
	return "", proto.LoginError_ErrPlatformSide
}
