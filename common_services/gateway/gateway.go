package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fananchong/gotcp"
)

type Gateway struct {
	srv  *gotcp.Server             // 面向客户端的TCP服务器
	sess map[string]*gotcp.Session // 做为客户端连接服务器组内服务的网络会话
}

func NewGateway() *Gateway {
	this := &Gateway{
		srv:  &gotcp.Server{},
		sess: make(map[string]*gotcp.Session),
	}
	return this
}

func (this *Gateway) Start() bool {
	gotcp.SetLogger(xlog)
	this.srv.RegisterSessType(UserSession{})
	addrinfo := strings.Split(xargs.ArgsBase.Pending.ExternalIp, ":")
	port, _ := strconv.Atoi(addrinfo[1])
	this.srv.Start(fmt.Sprintf(":%d", port))
	return true
}

func (this *Gateway) Close() {
	this.srv.Close()
}
