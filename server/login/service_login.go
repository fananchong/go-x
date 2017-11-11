package main

import (
	"net/http"

	"github.com/fananchong/go-x/common"
)

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
	this.InitHander()
	this.ListenAndServe(addr)
}

func (this *ServiceLogin) InitHander() {
	this.HandleFunc("/login", this.RequestLogin)
}

func (this *ServiceLogin) RequestLogin(w http.ResponseWriter, req *http.Request) {
	xlog.Infoln("hello")
}
