package main

import (
	"sync"

	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common_services/db"
	"github.com/fananchong/go-x/common_services/proto"
	"github.com/fananchong/gotcp"
)

type SessionAccount struct {
	gotcp.Session
	uid uint64
}

func (this *SessionAccount) OnRecv(data []byte, flag byte) {
	if this.IsVerified() == false {
		this.doVerify(data, flag)
		return
	}

	cmd := proto.MsgTypeCmd(gotcp.GetCmd(data))
	switch cmd {
	case proto.MsgTypeCmd_ForwardC:
		msg := &proto.MsgForwardC{}
		if gotcp.DecodeCmd(data, flag, msg) == nil {
			xlog.Debugln("decodeMsg fail.")
			this.Close()
			return
		}
		Forward(common.ServerType(msg.GetType()), msg.GetData())
	default:
		xlog.Debugln("unkown msg")
		this.Close()
	}
}

func (this *SessionAccount) OnClose() {
	if _, loaded := xaccounts.Load(this.uid); loaded {
		xaccounts.Delete(this.uid)
	}
}

func (this *SessionAccount) doVerify(data []byte, flag byte) {
	msg := &proto.MsgVerify{}
	if gotcp.DecodeCmd(data, flag, msg) == nil {
		xlog.Debugln("decodeMsg fail.")
		this.Close()
		return
	}

	token := db.NewToken(common.GetArgs().DbToken.Name, msg.GetAccount())
	if err := token.Load(); err != nil {
		xlog.Debugln(err)
		this.Close()
		return
	}

	if token.GetToken() != msg.GetToken() {
		xlog.Debugln("token error.")
		this.Close()
		return
	}

	uidserver := db.NewUIDServer(common.GetArgs().DbServer.Name, token.GetUid())
	uidserver.SetGateway(xnode.Id())
	if err := uidserver.Save(); err != nil {
		xlog.Debugln(err)
		this.Close()
		return
	}

	if s, loaded := xaccounts.Load(token.GetUid()); loaded {
		s.(*SessionAccount).uid = 0
		s.(*SessionAccount).Close()
	}
	xaccounts.Store(token.GetUid(), this)
	this.uid = token.GetUid()

	kickmsg := &proto.MsgKick{}
	kickmsg.UID = this.uid
	ForwardMsg(common.Hub, proto.MsgTypeCmd_Kick, kickmsg)

	this.Verify()

	rep := &proto.MsgVerifySuccess{}
	this.SendMsg(uint64(proto.MsgTypeCmd_VerifySuccess), rep)

	xlog.Debugln("account:", msg.GetAccount(), "verify success.")
}

// 由于Gateway功能相当简单，这里session管理，没有做成单例管理类。
// 请不要模仿这种不好的习惯:)
var xaccounts sync.Map

func initRedis() bool {
	go_redis_orm.SetNewRedisHandler(go_redis_orm.NewDefaultRedisClient)
	dbTokenName := common.GetArgs().DbToken.Name
	err := go_redis_orm.CreateDB(dbTokenName, common.GetArgs().DbToken.Addrs, common.GetArgs().DbToken.Password, common.GetArgs().DbToken.DBIndex)
	if err != nil {
		common.GetLogger().Errorln(err)
		return false
	}
	dbTokenName = common.GetArgs().DbServer.Name
	err = go_redis_orm.CreateDB(dbTokenName, common.GetArgs().DbServer.Addrs, common.GetArgs().DbServer.Password, common.GetArgs().DbServer.DBIndex)
	if err != nil {
		common.GetLogger().Errorln(err)
		return false
	}
	return true
}
