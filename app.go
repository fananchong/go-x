package internal

import (
	"github.com/fananchong/go-x/internal/common"
)

type App struct {
	common.App
}

func NewApp() *App {
	this := &App{}
	this.Type = int(common.Login)
	this.Args = XARGS
	this.Logger = XLOG
	this.Derived = this
	return this
}

var runner = NewLogin()

func (this *App) OnAppReady() {
	go func() {
		if runner.Start() == false {
			this.Close()
		}
	}()
}

func (this *App) OnAppShutDown() {
	runner.Close()
}
