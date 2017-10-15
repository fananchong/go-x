package main

import (
	"common"
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
