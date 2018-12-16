package common

import (
	"fmt"

	"github.com/fananchong/go-x/base"
	discovery "github.com/fananchong/go-x/internal/common/k8s/serverlist"
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
	gotcp.SetLogger(base.XLOG)
	if this.startServer() == false {
		return false
	}
	return true
}

func (this *TcpServer) startServer() bool {
	port := discovery.GetNode().Ports("")
	return this.Server.Start(fmt.Sprintf(":%d", port))
}

func (this *TcpServer) Close() {
	this.Server.Close()
}
