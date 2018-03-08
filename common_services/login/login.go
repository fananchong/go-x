package main

import (
	"net/http"

	pb "github.com/fananchong/go-proto-helper"
	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/db"
	"github.com/fananchong/go-x/common/proto"
)

type LoginMsgHandlerType func(http.ResponseWriter, *http.Request, string, string)

type Login struct {
	common.WebService
	cmds          map[proto.MsgTypeCmd]LoginMsgHandlerType
	dbAccountName string
	dbTokenName   string
	suid          *db.SUID
}

func NewLogin() *Login {
	return &Login{}
}

func (this *Login) Start() bool {
	if this.cmds == nil {
		this.cmds = make(map[proto.MsgTypeCmd]LoginMsgHandlerType)
		this.cmds[proto.MsgTypeCmd_Login] = this.MsgLogin
	}

	go_redis_orm.SetNewRedisHandler(go_redis_orm.NewDefaultRedisClient)

	// db account
	this.dbAccountName = common.GetArgs().DbAccount.Name
	err := go_redis_orm.CreateDB(this.dbAccountName, common.GetArgs().DbAccount.Addrs, common.GetArgs().DbAccount.Password, common.GetArgs().DbAccount.DBIndex)
	if err != nil {
		common.GetLogger().Errorln(err)
		return false
	}

	// db token
	this.dbTokenName = common.GetArgs().DbToken.Name
	err = go_redis_orm.CreateDB(this.dbTokenName, common.GetArgs().DbToken.Addrs, common.GetArgs().DbToken.Password, common.GetArgs().DbToken.DBIndex)
	if err != nil {
		common.GetLogger().Errorln(err)
		return false
	}

	// suid
	this.suid = &db.SUID{Cli: go_redis_orm.GetDB(this.dbAccountName)}

	// logger
	pb.SetLogger(common.GetLogger())

	// http service
	this.HandleFunc("/msg", this.request)
	this.ListenAndServe(xargs.Login.Listen)
	return true
}

func (this *Login) Register(cmd proto.MsgTypeCmd, f LoginMsgHandlerType) {
	if _, ok := this.cmds[cmd]; !ok {
		this.cmds[cmd] = f
	} else {
		panic("Register fail.")
	}
}
