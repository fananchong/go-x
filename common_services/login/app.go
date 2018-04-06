package main

import (
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/discovery"
)

var (
	xargs  *Args           = NewArgs()
	xlog   common.ILogger  = common.NewGLogger()
	xnode  *discovery.Node = &discovery.Node{}
	xlogin *Login          = NewLogin()
	xapp   *App            = NewApp()
)

type App struct {
	common.App
}

func NewApp() *App {
	this := &App{}
	this.Type = int(common.Login)
	this.Derived = this
	this.Args = xargs
	this.Logger = xlog
	this.Node = xnode
	return this
}

func (this *App) OnAppReady() {
	go func() {
		if xlogin.Start() == false {
			this.Close()
		}
	}()
}

func (this *App) OnAppShutDown() {
}
