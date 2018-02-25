package godiscovery

import (
	"context"
	"runtime/debug"
	"strconv"
	"sync"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

type IWatch interface {
	INode
	OnNodeUpdate(nodeType int, id string, data []byte)
	OnNodeJoin(nodeType int, id string, data []byte)
	OnNodeLeave(nodeType int, id string)
}

type Watch struct {
	Derived IWatch
	nodes   map[int]map[string]int
	mutex   sync.Mutex
}

func (this *Watch) Open(watchNodeTypes []int) {
	this.nodes = make(map[int]map[string]int)
	for _, nodeType := range watchNodeTypes {
		this.nodes[nodeType] = make(map[string]int)
		go this.watch(nodeType)
	}
}

func (this *Watch) watch(nodeType int) {
	this.Derived.GetLogger().Infoln("start watch node, node type =", nodeType)
	defer func() {
		if err := recover(); err != nil {
			this.Derived.GetLogger().Errorln("[异常] ", err, "\n", string(debug.Stack()))
		}
		this.Derived.Close()
	}()
	prefix := strconv.Itoa(nodeType) + "-"
	rch := this.Derived.GetClient().Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			key := string(ev.Kv.Key)
			if key == this.Derived.Id() {
				continue
			}
			if ev.Type == mvccpb.PUT {
				this.mutex.Lock()
				if _, ok := this.nodes[nodeType][key]; ok {
					this.mutex.Unlock()
					this.Derived.OnNodeUpdate(nodeType, key, ev.Kv.Value)
				} else {
					this.nodes[nodeType][key] = 1
					this.mutex.Unlock()
					this.Derived.OnNodeJoin(nodeType, key, ev.Kv.Value)
				}
			} else if ev.Type == mvccpb.DELETE {
				this.mutex.Lock()
				if _, ok := this.nodes[nodeType][key]; ok {
					delete(this.nodes[nodeType], key)
					this.mutex.Unlock()
					this.Derived.OnNodeLeave(nodeType, key)
				} else {
					this.mutex.Unlock()
				}
			} else {
				panic("unknow error!")
			}
		}
	}
}

func (this *Watch) Close() {
}
