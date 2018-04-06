package main

import (
	"sync"

	"github.com/fananchong/go-x/common/discovery"
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

func (this *Node) OnNodeUpdate(nodeType int, id string, data []byte) {
	this.Node.OnNodeUpdate(nodeType, id, data)
	if this.has(id) == false {
		info, ok := this.Servers.GetByID(id)
		if !ok {
			xlog.Errorln("can't find server info")
			return
		}
		this.tryConnect(id, info)
	}
}

func (this *Node) OnNodeJoin(nodeType int, id string, data []byte) {
	this.Node.OnNodeJoin(nodeType, id, data)
	info, ok := this.Servers.GetByID(id)
	if !ok {
		xlog.Errorln("can't find server info")
		return
	}
	this.tryConnect(id, info)
}

func (this *Node) OnNodeLeave(nodeType int, id string) {
	this.Node.OnNodeLeave(nodeType, id)
	this.tryDelete(id)
}

func (this *Node) tryConnect(id string, info *discovery.ServerInfo) {
	this.tryDelete(id)
	session := &SessionNode{}
	this.pending.Store(id, session)
	go func() {
		if session.Connect(info.ExternalIp, session) == true {
			this.nodes.Store(id, session)
		}
		this.pending.Delete(id)
	}()
}

func (this *Node) tryDelete(id string) {
	if session, loaded := this.nodes.Load(id); loaded {
		session.(*SessionNode).Close()
		this.nodes.Delete(id)
	}
	if session, loaded := this.pending.Load(id); loaded {
		session.(*SessionNode).Close()
		this.pending.Delete(id)
	}
}

func (this *Node) has(id string) bool {
	if _, loaded := this.nodes.Load(id); loaded {
		return true
	}
	if _, loaded := this.pending.Load(id); loaded {
		return true
	}
	return false
}
