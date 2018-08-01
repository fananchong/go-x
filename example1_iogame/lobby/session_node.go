package main

import (
	"context"
	"net"

	service "github.com/fananchong/go-x/common_services"
	"github.com/fananchong/go-x/example1_iogame/proto"
	"github.com/fananchong/gotcp"
)

type SessionNode struct {
	service.SessionIntranet
}

func (this *SessionNode) Init(conn net.Conn, root context.Context, derived gotcp.ISession) {
	this.SessionIntranet.Init(conn, root, derived)

	// init cmd
	this.SessionIntranet.Msgs.Handlers[uint32(proto.MsgTypeCmd_Lobby_PlayerBaseInfo)] = this.cmdPlayerBaseInfo
}

func (this *SessionNode) cmdPlayerBaseInfo(uid uint64, data []byte, flag byte) {
	xlog.Debugln("cmdPlayerBaseInfo")
	rep := &proto.MsgPlayerBaseInfoResult{}
	rep.Name = "test1" // TODO: 先调通网络消息流程，再处理角色数据相关
	this.SendMsgtoClient(uid, uint64(proto.MsgTypeCmd_Lobby_PlayerBaseInfo), rep)
}
