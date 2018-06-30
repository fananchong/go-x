package main

import (
	"sync"

	discovery "github.com/fananchong/go-x/common/k8s/serverlist"
	"github.com/fananchong/go-x/common_services/proto"
)

type Node struct {
	discovery.Node
	nodes   sync.Map
	pending sync.Map
}

func NewNode() *Node {
	this := &Node{}
	return this
}

func (this *Node) OnNodeUpdate(nodeIp string, nodeType int, id uint32, data []byte) {
	this.Node.OnNodeUpdate(nodeIp, nodeType, id, data)
	if this.has(id) == false {
		info, ok := this.Servers.GetByID(id)
		if !ok {
			xlog.Errorln("can't find server info")
			return
		}
		this.tryConnect(id, info)
	}
}

func (this *Node) OnNodeJoin(nodeIp string, nodeType int, id uint32, data []byte) {
	this.Node.OnNodeJoin(nodeIp, nodeType, id, data)
	info, ok := this.Servers.GetByID(id)
	if !ok {
		xlog.Errorln("can't find server info")
		return
	}
	this.tryConnect(id, info)
}

func (this *Node) OnNodeLeave(nodeType int, id uint32) {
	this.Node.OnNodeLeave(nodeType, id)
	this.tryDelete(id)
}

func (this *Node) tryConnect(id uint32, info *discovery.ServerInfo) {
	this.tryDelete(id)
	session := &SessionNode{}
	this.pending.Store(id, session)
	go func() {
		if session.Connect(info.ExternalIp, session) == true {
			this.nodes.Store(id, session)
			session.id = id
			session.t = info.GetType()

			msg := &proto.MsgVerifyS{}
			msg.Id = this.Node.Id()
			msg.Token = xargs.Common.IntranetToken
			session.SendMsg(uint64(proto.MsgTypeCmd_Verify), msg)
		}
		this.pending.Delete(id)
	}()
}

func (this *Node) tryDelete(id uint32) {
	if session, loaded := this.nodes.Load(id); loaded {
		session.(*SessionNode).Close()
		this.nodes.Delete(id)
	}
	if session, loaded := this.pending.Load(id); loaded {
		session.(*SessionNode).Close()
		this.pending.Delete(id)
	}
}

func (this *Node) has(id uint32) bool {
	if _, loaded := this.nodes.Load(id); loaded {
		return true
	}
	if _, loaded := this.pending.Load(id); loaded {
		return true
	}
	return false
}
