package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fananchong/go-x/base"
	"github.com/fananchong/go-x/common_services/proto"
	"github.com/fananchong/go-x/internal/k8s"
	discovery "github.com/fananchong/go-x/internal/k8s/serverlist"
)

type Node struct {
	*discovery.Node
	nodes   map[uint32]*SessionNode
	pending map[uint32]*SessionNode
	mutex   sync.Mutex
}

func NewNode() *Node {
	this := &Node{
		nodes:   make(map[uint32]*SessionNode),
		pending: make(map[uint32]*SessionNode),
	}
	this.Node = discovery.NewNode()
	go this.loopCheck()
	return this
}

func (this *Node) OnNodeJoin(endpoint *k8s.Endpoint) {
	this.Node.OnNodeJoin(endpoint)
	id := endpoint.Id()
	defer this.mutex.Unlock()
	this.mutex.Lock()
	this.tryConnect(id, endpoint)
}

func (this *Node) OnNodeLeave(endpoint *k8s.Endpoint) {
	this.Node.OnNodeLeave(endpoint)
	id := endpoint.Id()
	defer this.mutex.Unlock()
	this.mutex.Lock()
	this.tryDelete(id)
}

func (this *Node) tryConnect(id uint32, endpoint *k8s.Endpoint) {
	this.tryDelete(id)
	session := &SessionNode{}
	this.pending[id] = session
	session.endpoint = endpoint
}

func (this *Node) tryConnectDetail(id uint32, session *SessionNode) bool {
	addr := fmt.Sprintf("%s:%d", session.endpoint.IP, session.endpoint.Ports[""])
	if session.Connect(addr, session) == true {
		this.nodes[id] = session
		msg := &proto.MsgVerifyS{}
		msg.Id = this.Node.Id()
		msg.Token = base.XARGS.Common.IntranetToken
		session.SendMsg(uint64(proto.MsgTypeCmd_Verify), msg)
		return true
	}
	return false
}

func (this *Node) tryDelete(id uint32) {
	if session, ok := this.nodes[id]; ok {
		session.Close()
		delete(this.nodes, id)
	}
	if session, ok := this.pending[id]; ok {
		session.Close()
		delete(this.pending, id)
	}
}

func (this *Node) loopCheck() {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			this.mutex.Lock()
			var dels []uint32
			for id, session := range this.pending {
				if this.tryConnectDetail(id, session) {
					dels = append(dels, id)
				}
			}
			for _, id := range dels {
				delete(this.pending, id)
			}
			this.mutex.Unlock()
		}
	}
}
