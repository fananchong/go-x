package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fananchong/gotcp"
)

type Lobby struct {
	srv *gotcp.Server
}

func NewLobby() *Lobby {
	this := &Lobby{
		srv: &gotcp.Server{},
	}
	return this
}

func (this *Lobby) Start() bool {
	gotcp.SetLogger(xlog)
	if this.startServer() == false {
		return false
	}
	return true
}

func (this *Lobby) startServer() bool {
	this.srv.RegisterSessType(SessionNode{})
	addrinfo := strings.Split(xargs.ArgsBase.Pending.ExternalIp, ":")
	port, _ := strconv.Atoi(addrinfo[1])
	return this.srv.Start(fmt.Sprintf(":%d", port))
}

func (this *Lobby) Close() {
	this.srv.Close()
}
