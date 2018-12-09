package main

import (
	"context"
	"net"

	service "github.com/fananchong/go-x/common_services"
	"github.com/fananchong/gotcp"
)

type SessionNode struct {
	service.SessionIntranet
}

func (this *SessionNode) Init(conn net.Conn, root context.Context, derived gotcp.ISession) {
	this.SessionIntranet.Init(conn, root, derived)

	// init cmd
	this.DefaultHandler = this.cmdDefault
}

func (this *SessionNode) cmdDefault(_ uint64, data []byte, flag byte) {
	this.BroadcastExcludeMe(data, flag)
}
