package main

import "github.com/fananchong/go-x/common/login"

var (
	xlogin *ServiceLogin = NewServiceLogin()
)

type ServiceLogin struct {
	login.Login
}

func NewServiceLogin() *ServiceLogin {
	return &ServiceLogin{}
}
