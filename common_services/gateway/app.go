package main

import (
	"github.com/fananchong/go-redis-orm.v2"
	"github.com/fananchong/go-x/common"
)

var (
	xargs *Args          = &Args{}
	xlog  common.ILogger = common.NewGLogger()
	xnode *Node          = NewNode()
)

type App struct {
	common.App
}

func NewApp() *App {
	this := &App{}
	this.Type = int(common.Gateway)
	this.Args = xargs
	this.Logger = xlog
	this.Node = xnode
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
		runner.RegisterSessType(SessionAccount{})
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
		common.GetArgs().DbToken.Name,
		common.GetArgs().DbToken.Addrs,
		common.GetArgs().DbToken.Password,
		common.GetArgs().DbToken.DBIndex)
	if err != nil {
		common.GetLogger().Errorln(err)
		return false
	}
	err = go_redis_orm.CreateDB(
		common.GetArgs().DbServer.Name,
		common.GetArgs().DbServer.Addrs,
		common.GetArgs().DbServer.Password,
		common.GetArgs().DbServer.DBIndex)
	if err != nil {
		common.GetLogger().Errorln(err)
		return false
	}
	return true
}
