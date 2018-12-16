package main

import (
	"net/http"

	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/fananchong/go-x/base"
	"github.com/fananchong/go-x/common_services/db"
	"github.com/fananchong/go-x/common_services/proto"
	"github.com/fananchong/gotcp"
)

type LoginMsgHandlerType func(http.ResponseWriter, *http.Request, string, string)

type Login struct {
	base.WebService
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
	this.dbAccountName = base.XARGS.DbAccount.Name
	err := go_redis_orm.CreateDB(this.dbAccountName, base.XARGS.DbAccount.Addrs, base.XARGS.DbAccount.Password, base.XARGS.DbAccount.DBIndex)
	if err != nil {
		base.XLOG.Errorln(err)
		return false
	}

	// db token
	this.dbTokenName = base.XARGS.DbToken.Name
	err = go_redis_orm.CreateDB(this.dbTokenName, base.XARGS.DbToken.Addrs, base.XARGS.DbToken.Password, base.XARGS.DbToken.DBIndex)
	if err != nil {
		base.XLOG.Errorln(err)
		return false
	}

	// suid
	this.suid = &db.SUID{Cli: go_redis_orm.GetDB(this.dbAccountName)}

	// logger
	gotcp.SetLogger(base.XLOG)

	// http service
	this.HandleFunc("/msg", this.request)
	this.ListenAndServe(externArgs.Listen)
	return true
}

func (this *Login) Register(cmd proto.MsgTypeCmd, f LoginMsgHandlerType) {
	if _, ok := this.cmds[cmd]; !ok {
		this.cmds[cmd] = f
	} else {
		panic("Register fail.")
	}
}

func (this *Login) Close() {
	this.WebService.Close()
}
