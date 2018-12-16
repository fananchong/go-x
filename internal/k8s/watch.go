package k8s

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Watch struct {
	nt        *NodeType
	endpoints sync.Map
	ctx       context.Context
	ctxCancel context.CancelFunc
	tick      *time.Ticker
	derived   INode
}

func NewWatch(nt *NodeType, ctx context.Context, d time.Duration, derived INode) *Watch {
	this := &Watch{
		nt:      nt,
		derived: derived,
	}
	this.init(ctx, d)
	return this
}

func (this *Watch) init(ctx context.Context, d time.Duration) {
	this.ctx, this.ctxCancel = context.WithCancel(ctx)
	this.tick = time.NewTicker(d)
	go func() {
		for {
			select {
			case <-this.tick.C:
				this.checkEndpoints()
			case <-this.ctx.Done():
				this.close()
				return
			}
		}
	}()
}

func (this *Watch) checkEndpoints() {
	eps, err := GetEndpoints(this.nt.ns, this.nt.svc)
	if err == nil {
		for _, ep := range eps {
			if _, ok := this.endpoints.Load(ep.Index); !ok {
				ep.NodeType = this.nt.t
				this.endpoints.Store(ep.Index, ep)
				this.derived.OnNodeJoin(ep)
			}
		}
	} else {
		fmt.Println(err)
	}
}

func (this *Watch) close() {
	if this.tick != nil {
		this.tick.Stop()
		this.tick = nil
	}
}

func (this *Watch) OnLoseEndpoint(index int) {
	this.endpoints.Delete(index)
}
