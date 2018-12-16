package main

import (
	"github.com/fananchong/go-x/base"
	"github.com/fananchong/go-x/internal"
)

type App struct {
	internal.App
	runner base.Plugin
}

func NewApp(runner base.Plugin) *App {
	this := &App{}
	this.Derived = this
	this.runner = runner
	return this
}

func (this *App) OnAppReady() {
	if this.runner.Init() == false {
		this.Close()
		return
	}
	go func() {
		if this.runner.Start() == false {
			this.Close()
		}
	}()
}

func (this *App) OnAppShutDown() {
	this.runner.Close()
}
