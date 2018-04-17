package main

import (
	"sync"

	"github.com/fananchong/go-x/common_services/proto"
	"github.com/fananchong/gotcp"
)

type SessionNode struct {
	gotcp.Session
	id string
}

func (this *SessionNode) OnRecv(data []byte, flag byte) {
	if this.IsVerified() == false {
		this.doVerify(data, flag)
		return
	}
	cmd := proto.MsgTypeCmd(gotcp.GetCmd(data))
	switch cmd {
	default:
		xnodes.Range(func(key interface{}, val interface{}) bool {
			if key.(string) != this.id {
				val.(*SessionNode).Send(data, flag)
			}
			return true
		})
	}
}

func (this *SessionNode) doVerify(data []byte, flag byte) {
	msg := &proto.MsgVerify{}
	if gotcp.DecodeCmd(data, flag, msg) == nil {
		xlog.Errorln("decodeMsg fail.")
		this.Close()
		return
	}
	if msg.GetToken() != xargs.Common.IntranetToken {
		xlog.Errorln("token error.")
		this.Close()
	}
	this.id = msg.GetAccount()
	xnodes.Store(this.id, this)
	this.Verify()
	xlog.Debugln("id:", msg.GetAccount(), "verify success.")
}

func (this *SessionNode) OnClose() {
	if _, loaded := xnodes.Load(this.id); loaded {
		xnodes.Delete(this.id)
	}
}

var xnodes sync.Map
