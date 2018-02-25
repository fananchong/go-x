package main

import (
	"github.com/fananchong/go-x/common/login"
	"github.com/fananchong/go-x/common/proto"
)

var (
	xlogin *ServiceLogin = NewServiceLogin()
)

type ServiceLogin struct {
	login.Login
}

func NewServiceLogin() *ServiceLogin {
	this := &ServiceLogin{}
	this.Derived = this
	return this
}

// 第3方平台，获取帐号密码
func (this *ServiceLogin) GetPassword(account string, mode proto.LoginMode, userdata []byte) (string, proto.LoginError) {
	return "", proto.LoginError_ErrPlatformSide
}
