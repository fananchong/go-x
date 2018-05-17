package common

import (
	"fmt"

	discovery "github.com/fananchong/go-discovery/serverlist"
	"github.com/fananchong/gotcp"
)

type TcpServer struct {
	*gotcp.Server
}

func NewTcpServer() *TcpServer {
	this := &TcpServer{
		Server: &gotcp.Server{},
	}
	return this
}

func (this *TcpServer) Start() bool {
	gotcp.SetLogger(xlog)
	if this.startServer() == false {
		return false
	}
	return true
}

func (this *TcpServer) startServer() bool {
	port := discovery.GetNode().GetPort()
	return this.Server.Start(fmt.Sprintf(":%d", port))
}

func (this *TcpServer) Close() {
	this.Server.Close()
}
