package main

import (
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/example1_iogame"
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
	this.Type = int(iogame.Room)
	this.Args = xargs
	this.Logger = xlog
	this.Derived = this
	return this
}

func (this *App) OnAppReady() {
}

func (this *App) OnAppShutDown() {
}
