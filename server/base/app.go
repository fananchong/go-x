package main

import (
	"github.com/fananchong/go-x/common"
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
	initNode()
}

func (this *App) OnAppShutDown() {

}

func initNode() {
	node := NewNode()
	node.SetLogger(xlog)
	node.OpenByStr(xargs.EtcdHosts, int(xargs.NodeType), xargs.WatchNodeTypes, xargs.PutInterval)
}
