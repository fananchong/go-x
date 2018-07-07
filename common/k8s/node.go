package k8s

import (
	"context"
	"errors"
	"os"
	"time"
)

type INode interface {
	OnNodeJoin(endpoint *Endpoint)
	OnNodeLeave(endpoint *Endpoint)
	GetBase() interface{}
}

type Node struct {
	endpoint  *Endpoint
	ctx       context.Context
	ctxCancel context.CancelFunc
	watchs    []*Watch
}

func (this *Node) Init(nodeType int, watchNodeTypes []int, d time.Duration, inst INode) error {
	if nt, ok := gNodeTypes[nodeType]; ok {
		this.ctx, this.ctxCancel = context.WithCancel(context.Background())
		index := getIndex(nt.svc, os.Getenv("POD_NAME"))
	LABEL_GETEP:
		if eps, err := GetEndpoints(nt.ns, nt.svc); err == nil {
			for _, ep := range eps {
				if index == ep.Index {
					this.endpoint = ep
					break
				}
			}
			if this.endpoint == nil {
				time.Sleep(1 * time.Second)
				goto LABEL_GETEP
			}
			for _, v := range watchNodeTypes {
				if nt, ok := gNodeTypes[v]; ok {
					w := NewWatch(nt, this.ctx, d)
					w.Derived = inst
					this.watchs = append(this.watchs, w)
				}
			}
			return nil
		} else {
			return err
		}
	} else {
		return errors.New("no find nodeType info!")
	}
}

func (this *Node) Close() {
	if this.ctxCancel != nil {
		this.ctxCancel()
		this.ctxCancel = nil
	}
}

func (this *Node) Id() uint32 {
	return this.endpoint.Id()
}

func (this *Node) Ports(name string) int {
	return this.endpoint.Ports[name]
}

func (this *Node) GetBase() interface{} {
	return this
}

// 子类可以根据需要重载下面的方法
//     注意 OnNodeJoin 、 OnNodeLeave 在内部协程被调用，请注意多协程安全！！！
func (this *Node) OnNodeJoin(endpoint *Endpoint) {

}

func (this *Node) OnNodeLeave(endpoint *Endpoint) {
	for _, v := range this.watchs {
		v.OnLoseEndpoint(endpoint.Index)
	}
}

