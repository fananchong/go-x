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

func (this *ServiceLogin) OnVerifyAccount(string, string, proto.LoginMode, []byte) (uint64, error) {
	return 0, nil
}
