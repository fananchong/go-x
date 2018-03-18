package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/discovery"
	"github.com/fananchong/gotcp"
)

var (
	xargs *Args                     = NewArgs()
	xlog  common.ILogger            = common.NewGLogger()
	xnode *discovery.Node           = &discovery.Node{}
	xapp  *App                      = NewApp()
	xsrv  *gotcp.Server             = &gotcp.Server{}                 // 面向客户端的TCP服务器
	xsess map[string]*gotcp.Session = make(map[string]*gotcp.Session) // 做为客户端连接服务器组内服务的网络会话
)

type App struct {
	common.App
}

func NewApp() *App {
	this := &App{}
	this.Derived = this
	this.Args = xargs
	this.Logger = xlog
	this.Node = xnode
	return this
}

func (this *App) OnAppReady() {
	this.startUserServer()
}

func (this *App) OnAppShutDown() {
	xsrv.Close()
}

func (this *App) startUserServer() {
	gotcp.SetLogger(xlog)
	xsrv.RegisterSessType(UserSession{})
	addrinfo := strings.Split(xargs.Gateway.ExternalIp, ":")
	var port int
	var err error
	if len(addrinfo) < 2 {
		port = gotcp.GetVaildPort()
	} else {
		port, err = strconv.Atoi(addrinfo[1])
		if err != nil {
			panic(err)
			return
		}
	}
	xsrv.Start(fmt.Sprintf(":%d", port))
}
