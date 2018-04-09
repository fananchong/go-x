package common

import (
	"fmt"
	"strconv"
	"strings"

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
	addrinfo := strings.Split(xargs.Pending.ExternalIp, ":")
	port, _ := strconv.Atoi(addrinfo[1])
	return this.Server.Start(fmt.Sprintf(":%d", port))
}

func (this *TcpServer) Close() {
	this.Server.Close()
}
