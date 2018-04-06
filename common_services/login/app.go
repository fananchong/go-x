package main

import (
	"github.com/fananchong/go-x/common"
)

var (
	xargs  *Args          = &Args{}
	xlog   common.ILogger = common.NewGLogger()
	xlogin *Login         = NewLogin()
	xapp   *App           = NewApp()
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
