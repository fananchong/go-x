package main

import (
	"sync"

	"github.com/fananchong/gotcp"
)

type SessionNode struct {
	gotcp.Session
	id string
	t  uint32
}

func (this *SessionNode) OnRecv(data []byte, flag byte) {
	if this.IsVerified() == false {
		xnodes.Store(this.id, this)
		xnodesMutex.Lock()
		defer xnodesMutex.Unlock()
		if _, ok := xnodesByType[this.t]; !ok {
			xnodesByType[this.t] = make(map[string]*SessionNode)
		}
		xnodesByType[this.t][this.id] = this
		this.Verify()
	}
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
var xnodesByType map[uint32]map[string]*SessionNode
var xnodesMutex sync.RWMutex
