package login

import (
	"net/http"

	pb "github.com/fananchong/go-proto-helper"
	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/db"
	"github.com/fananchong/go-x/common/proto"
)

type ILogin interface {
	GetPassword(string, proto.LoginMode, []byte) (string, proto.LoginError)
}

type LoginMsgHandlerType func(http.ResponseWriter, *http.Request, string, string)

type Login struct {
	common.WebService
	cmds    map[proto.MsgTypeCmd]LoginMsgHandlerType
	dbName  string
	suid    *db.SUID
	Derived ILogin
}

func (this *Login) Start() bool {
	if this.cmds == nil {
		this.cmds = make(map[proto.MsgTypeCmd]LoginMsgHandlerType)
		this.cmds[proto.MsgTypeCmd_Login] = this.MsgLogin
	}

	this.dbName = common.GetArgs().DbAccount.Name
	go_redis_orm.SetNewRedisHandler(go_redis_orm.NewDefaultRedisClient)
	err := go_redis_orm.CreateDB(this.dbName, common.GetArgs().DbAccount.Addrs, common.GetArgs().DbAccount.Password, common.GetArgs().DbAccount.DBIndex)
	if err != nil {
		common.GetLogger().Errorln(err)
		return false
	}
	this.suid = &db.SUID{Cli: go_redis_orm.GetDB(this.dbName)}
	pb.SetLogger(common.GetLogger())
	this.HandleFunc("/msg", this.request)
	this.ListenAndServe(common.GetArgs().Login.Listen)
	return true
}

func (this *Login) Register(cmd proto.MsgTypeCmd, f LoginMsgHandlerType) {
	if _, ok := this.cmds[cmd]; !ok {
		this.cmds[cmd] = f
	} else {
		panic("Register fail.")
	}
}
