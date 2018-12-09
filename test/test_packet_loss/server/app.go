package main

import (
	"github.com/fananchong/go-x/common"
)

var (
	xlog common.ILogger = common.NewDefaultLogger()
)

type App struct {
	common.App
}

func NewApp() *App {
	this := &App{}
	this.Derived = this
	return this
}

func (this *App) OnAppReady() {
	go TcpServer()
	go UdpServer()
}

func (this *App) OnAppShutDown() {

}
