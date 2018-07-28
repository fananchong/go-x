package main

import (
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/example1_iogame/src"
)

var (
	xargs *Args          = &Args{}
	xlog  common.ILogger = common.NewGLogger()
)

type App struct {
	common.App
}

func NewApp() *App {
	this := &App{}
	this.Type = int(iogame.Lobby)
	this.Args = xargs
	this.Logger = xlog
	this.Derived = this
	return this
}

var runner = common.NewTcpServer()

func (this *App) OnAppReady() {
	go func() {
		runner.RegisterSessType(SessionNode{})
		if runner.Start() == false {
			this.Close()
		}
	}()
}

func (this *App) OnAppShutDown() {
	runner.Close()
}
