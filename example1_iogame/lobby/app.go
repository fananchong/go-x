package main

import (
	"github.com/fananchong/go-redis-orm.v2"
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
	this.Type = int(iogame.Lobby)
	this.Args = xargs
	this.Logger = xlog
	this.Derived = this
	return this
}

var runner = common.NewTcpServer()

func (this *App) OnAppReady() {
	if initRedis() == false {
		this.Close()
		return
	}
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

func initRedis() bool {
	go_redis_orm.SetNewRedisHandler(go_redis_orm.NewDefaultRedisClient)
	err := go_redis_orm.CreateDB(
		common.GetArgs().DbAccount.Name,
		common.GetArgs().DbAccount.Addrs,
		common.GetArgs().DbAccount.Password,
		common.GetArgs().DbAccount.DBIndex)
	if err != nil {
		common.GetLogger().Errorln(err)
		return false
	}
	return true
}
