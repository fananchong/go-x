package main

import (
	"fmt"
	"sync"

	"github.com/fananchong/go-x/common/k8s"
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

func (this *Node) OnNodeJoin(endpoint *k8s.Endpoint) {
	this.Node.OnNodeJoin(endpoint)
	id := endpoint.Id()
	this.tryConnect(id, endpoint)
}

func (this *Node) OnNodeLeave(endpoint *k8s.Endpoint) {
	this.Node.OnNodeLeave(endpoint)
	id := endpoint.Id()
	this.tryDelete(id)
}

func (this *Node) tryConnect(id uint32, endpoint *k8s.Endpoint) {
	this.tryDelete(id)
	session := &SessionNode{}
	this.pending.Store(id, session)
	go func() {
		addr := fmt.Sprintf("%s:%d", endpoint.IP, endpoint.Ports[""])
		if session.Connect(addr, session) == true {
			this.nodes.Store(id, session)
			session.endpoint = endpoint

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
