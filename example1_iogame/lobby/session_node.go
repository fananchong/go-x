package main

import (
	"context"
	"net"

	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/fananchong/go-x/base"
	service "github.com/fananchong/go-x/common_services"
	"github.com/fananchong/go-x/example1_iogame/db"
	"github.com/fananchong/go-x/example1_iogame/proto"
	"github.com/fananchong/gotcp"
)

type SessionNode struct {
	service.SessionIntranet
}

func (this *SessionNode) Init(conn net.Conn, root context.Context, derived gotcp.ISession) {
	this.SessionIntranet.Init(conn, root, derived)

	// init cmd
	this.SessionIntranet.Msgs.Handlers[uint32(proto.MsgTypeCmd_Lobby_CreatePlayer)] = this.cmdCreatePlayer
	this.SessionIntranet.Msgs.Handlers[uint32(proto.MsgTypeCmd_Lobby_PlayerBaseInfo)] = this.cmdPlayerBaseInfo
}

func (this *SessionNode) cmdCreatePlayer(uid uint64, data []byte, flag byte) {
	base.XLOG.Debugln("cmdCreatePlayer, uid =", uid)

	msg := &proto.MsgCreatePlayer{}
	if gotcp.DecodeCmd(data, flag, msg) == nil {
		base.XLOG.Debugln("decodeMsg fail. MsgCreatePlayer, uid =", uid)
		this.Close()
		return
	}

	role := db.NewRole(DbAccount, uid)
	err := role.Load()
	if err == nil {
		// 角色已存在
		rep := &proto.MsgCreatePlayerResult{}
		rep.Err = proto.EnumCreatePlayer_ErrExist
		this.SendMsgtoClient(uid, uint64(proto.MsgTypeCmd_Lobby_CreatePlayer), rep)
	} else {
		if err == go_redis_orm.ERR_ISNOT_EXIST_KEY {
			// 创建角色

			// TODO: 检查角色名、性别等是否合法

			role.SetName(msg.GetName())
			role.SetSex(uint8(msg.GetSex()))
			if err := role.Save(); err != nil {
				// 数据库错误
				rep := &proto.MsgCreatePlayerResult{}
				rep.Err = proto.EnumCreatePlayer_ErrDB
				this.SendMsgtoClient(uid, uint64(proto.MsgTypeCmd_Lobby_CreatePlayer), rep)
				return
			}

			rep := &proto.MsgCreatePlayerResult{}
			rep.Err = proto.EnumCreatePlayer_NoErr
			this.SendMsgtoClient(uid, uint64(proto.MsgTypeCmd_Lobby_CreatePlayer), rep)
		} else {
			// 数据库错误
			rep := &proto.MsgCreatePlayerResult{}
			rep.Err = proto.EnumCreatePlayer_ErrDB
			this.SendMsgtoClient(uid, uint64(proto.MsgTypeCmd_Lobby_CreatePlayer), rep)
		}
	}
}

func (this *SessionNode) cmdPlayerBaseInfo(uid uint64, data []byte, flag byte) {
	base.XLOG.Debugln("cmdPlayerBaseInfo, uid =", uid)

	role := db.NewRole(DbAccount, uid)
	err := role.Load()
	if err == nil {
		rep := &proto.MsgPlayerBaseInfoResult{}
		rep.Err = proto.EnumPlayerBaseInfo_NoErr
		rep.Name = role.GetName()
		rep.Sex = int32(role.GetSex())
		this.SendMsgtoClient(uid, uint64(proto.MsgTypeCmd_Lobby_PlayerBaseInfo), rep)
	} else {
		if err == go_redis_orm.ERR_ISNOT_EXIST_KEY {
			rep := &proto.MsgPlayerBaseInfoResult{}
			rep.Err = proto.EnumPlayerBaseInfo_ErrDB
			this.SendMsgtoClient(uid, uint64(proto.MsgTypeCmd_Lobby_PlayerBaseInfo), rep)
		} else {
			rep := &proto.MsgPlayerBaseInfoResult{}
			rep.Err = proto.EnumPlayerBaseInfo_ErrNoExist
			this.SendMsgtoClient(uid, uint64(proto.MsgTypeCmd_Lobby_PlayerBaseInfo), rep)
		}
	}
}
