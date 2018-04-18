package main

import (
	"sync"

	"github.com/fananchong/go-x/common_services/proto"
	"github.com/fananchong/gotcp"
)

type SessionNode struct {
	gotcp.Session
	id string
	t  uint32
}

func (this *SessionNode) OnRecv(data []byte, flag byte) {
	if this.IsVerified() == false {
		this.doVerify(data, flag)
		return
	}
}

func (this *SessionNode) doVerify(data []byte, flag byte) {
	msg := &proto.MsgVerify{}
	if gotcp.DecodeCmd(data, flag, msg) == nil {
		xlog.Errorln("decodeMsg fail.")
		this.Close()
		return
	}

	if this.id != msg.GetAccount() {
		xlog.Errorln("verify fail. id error")
		this.Close()
		return
	}

	if msg.GetToken() != xargs.Common.IntranetToken {
		xlog.Errorln("verify fail. token error")
		this.Close()
		return
	}

	xnodes.Store(this.id, this)
	xnodesMutex.Lock()
	defer xnodesMutex.Unlock()
	if _, ok := xnodesByType[this.t]; !ok {
		xnodesByType[this.t] = make(map[string]*SessionNode)
	}
	xnodesByType[this.t][this.id] = this
	this.Verify()
	xlog.Debugln("Id:", msg.GetAccount(), "verify success.")
}

func (this *SessionNode) OnClose() {
	if _, loaded := xnodes.Load(this.id); loaded {
		xnodes.Delete(this.id)
	}
	xnodesMutex.Lock()
	defer xnodesMutex.Unlock()
	if items, ok := xnodesByType[this.t]; ok {
		if _, ok2 := items[this.id]; ok2 {
			delete(items, this.id)
		}
		if len(items) == 0 {
			delete(xnodesByType, this.t)
		}
	}
}

var xnodes sync.Map
var xnodesByType map[uint32]map[string]*SessionNode = make(map[uint32]map[string]*SessionNode)
var xnodesMutex sync.RWMutex
