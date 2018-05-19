package godiscovery

import (
	"context"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

type IWatch interface {
	INode
	OnNodeUpdate(nodeIP string, nodeType int, id uint32, data []byte)
	OnNodeJoin(nodeIP string, nodeType int, id uint32, data []byte)
	OnNodeLeave(nodeType int, id uint32)
}

type Watch struct {
	Derived   IWatch
	nodes     map[int]map[uint32]int
	mutex     sync.Mutex
	ctx       context.Context
	ctxCancel context.CancelFunc
}

func (this *Watch) Open(root context.Context, watchNodeTypes []int) {
	this.ctx, this.ctxCancel = context.WithCancel(root)
	this.nodes = make(map[int]map[uint32]int)
	for _, nodeType := range watchNodeTypes {
		this.mutex.Lock()
		this.nodes[nodeType] = make(map[uint32]int)
		this.mutex.Unlock()
		go this.watch(nodeType)
	}
}

func (this *Watch) watch(nodeType int) {
	xlog.Infoln("start watch node, node type =", nodeType)
	defer func() {
		if err := recover(); err != nil {
			xlog.Errorln("[except] ", err, "\n", string(debug.Stack()))
			if this.Derived.GetClient() != nil {
				go this.watch(nodeType)
			}
		}
	}()
	prefix := strconv.Itoa(nodeType) + "-"
	cli := this.Derived.GetClient()
	if cli == nil {
		return
	}
	rch := cli.Watch(this.ctx, prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			key := string(ev.Kv.Key)
			temp := strings.Split(key, "-")
			if len(temp) < 2 {
				continue
			}
			tempkey, err := strconv.Atoi(temp[1])
			if err != nil {
				continue
			}
			if uint32(tempkey) == this.Derived.Id() {
				continue
			}
			if ev.Type == mvccpb.PUT {
				temp := string(ev.Kv.Value)
				nodeIP := strings.Split(temp, "#")[0]
				data := ev.Kv.Value[len(nodeIP)+1:]
				this.mutex.Lock()
				if _, ok := this.nodes[nodeType][uint32(tempkey)]; ok {
					this.mutex.Unlock()
					this.Derived.OnNodeUpdate(nodeIP, nodeType, uint32(tempkey), data)
				} else {
					this.nodes[nodeType][uint32(tempkey)] = 1
					this.mutex.Unlock()
					this.Derived.OnNodeJoin(nodeIP, nodeType, uint32(tempkey), data)
				}
			} else if ev.Type == mvccpb.DELETE {
				this.mutex.Lock()
				if _, ok := this.nodes[nodeType][uint32(tempkey)]; ok {
					delete(this.nodes[nodeType], uint32(tempkey))
					this.mutex.Unlock()
					this.Derived.OnNodeLeave(nodeType, uint32(tempkey))
				} else {
					this.mutex.Unlock()
				}
			} else {
				panic("unknow error!")
			}
		}
	}
}
