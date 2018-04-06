package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fananchong/gotcp"
)

type Base struct {
	srv *gotcp.Server
}

func NewBase() *Base {
	this := &Base{
		srv: &gotcp.Server{},
	}
	return this
}

func (this *Base) Start() bool {
	gotcp.SetLogger(xlog)
	if this.startServer() == false {
		return false
	}
	return true
}

func (this *Base) startServer() bool {
	this.srv.RegisterSessType(SessionNode{})
	addrinfo := strings.Split(xargs.ArgsBase.Pending.ExternalIp, ":")
	port, _ := strconv.Atoi(addrinfo[1])
	return this.srv.Start(fmt.Sprintf(":%d", port))
}

func (this *Base) Close() {
	this.srv.Close()
}
