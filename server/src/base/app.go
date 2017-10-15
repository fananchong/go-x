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

}

func (this *App) OnAppShutDown() {

}
