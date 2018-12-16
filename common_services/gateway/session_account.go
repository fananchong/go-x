package main

import (
	"sync"

	"github.com/fananchong/go-x/base"
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
	if cmd <= proto.MsgTypeCmd_COMMON_CMD_END {
		base.XLOG.Debugln("inner msg!")
		this.Close()
	}
	serverType := int(cmd) / int(base.XARGS.Common.MsgCmdOffset)
	if serverType <= int(base.COMMON_SERVER_END) {
		base.XLOG.Debugln("inner server!")
		this.Close()
	}
	msg := &proto.MsgForward{}
	msg.UID = this.uid
	msg.Data = data
	msg.Flag = int32(flag)
	newData, newFlag, err := gotcp.EncodeCmd(uint64(proto.MsgTypeCmd_Forward), msg)
	if err != nil {
		base.XLOG.Errorln(err)
		this.Close()
	}
	Forward(serverType, newData, newFlag)
}

func (this *SessionAccount) OnClose() {
	if _, loaded := xaccounts.Load(this.uid); loaded {
		xaccounts.Delete(this.uid)
	}
}

func (this *SessionAccount) doVerify(data []byte, flag byte) {
	msg := &proto.MsgVerify{}
	if gotcp.DecodeCmd(data, flag, msg) == nil {
		base.XLOG.Debugln("decodeMsg fail.")
		this.Close()
		return
	}

	token := db.NewToken(base.XARGS.DbToken.Name, msg.GetAccount())
	if err := token.Load(); err != nil {
		base.XLOG.Debugln(err)
		this.Close()
		return
	}

	if token.GetToken() != msg.GetToken() {
		base.XLOG.Debugln("token error.")
		this.Close()
		return
	}

	uidserver := db.NewUIDServer(base.XARGS.DbServer.Name, token.GetUid())
	uidserver.SetGateway(XNODE.Id())
	if err := uidserver.Save(); err != nil {
		base.XLOG.Debugln(err)
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
	ForwardMsg(int(base.Mgr), proto.MsgTypeCmd_Kick, kickmsg)

	this.Verify()

	rep := &proto.MsgVerifySuccess{}
	this.SendMsg(uint64(proto.MsgTypeCmd_VerifySuccess), rep)

	base.XLOG.Debugln("account:", msg.GetAccount(), "verify success.")
}

// 由于Gateway功能相当简单，这里session管理，没有做成单例管理类。
// 请不要模仿这种不好的习惯:)
var xaccounts sync.Map
