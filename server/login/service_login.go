package main

import "github.com/fananchong/go-x/common"

var (
	xlogin *ServiceLogin = NewServiceLogin()
)

type ServiceLogin struct {
	common.WebService
}

func NewServiceLogin() *ServiceLogin {
	return &ServiceLogin{}
}

func (this *ServiceLogin) Start(addr string) {
	//
	this.ListenAndServe(addr)
}
