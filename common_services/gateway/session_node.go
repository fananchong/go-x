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

		// TODO: verify

		nodes.Store(this.id, this)

		mutex.Lock()
		defer mutex.Unlock()
		if _, ok := nodesByType[this.t]; !ok {
			nodesByType[this.t] = make(map[string]*SessionNode)
		}
		nodesByType[this.t][this.id] = this
		this.Verify()
	}
}

func (this *SessionNode) OnClose() {
	if _, loaded := nodes.Load(this.id); loaded {
		nodes.Delete(this.id)
	}
	mutex.Lock()
	defer mutex.Unlock()
	if items, ok := nodesByType[this.t]; ok {
		if _, ok2 := items[this.id]; ok2 {
			delete(items, this.id)
		}
		if len(items) == 0 {
			delete(nodesByType, this.t)
		}
	}
}

var nodes sync.Map
var nodesByType map[uint32]map[string]*SessionNode
var mutex sync.RWMutex
