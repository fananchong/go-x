package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fananchong/gotcp"
)

type Gateway struct {
	srv *gotcp.Server
}

func NewGateway() *Gateway {
	this := &Gateway{
		srv: &gotcp.Server{},
	}
	return this
}

func (this *Gateway) Start() bool {
	gotcp.SetLogger(xlog)
	if this.startServer() == false {
		return false
	}
	return true
}

func (this *Gateway) startServer() bool {
	this.srv.RegisterSessType(SessionAccount{})
	addrinfo := strings.Split(xargs.ArgsBase.Pending.ExternalIp, ":")
	port, _ := strconv.Atoi(addrinfo[1])
	return this.srv.Start(fmt.Sprintf(":%d", port))
}

func (this *Gateway) Close() {
	this.srv.Close()
}
