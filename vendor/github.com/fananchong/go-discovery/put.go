package godiscovery

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/coreos/etcd/clientv3"
	uuid "github.com/satori/go.uuid"
)

type IPut interface {
	INode
	GetPutData() (string, error)
}

type Put struct {
	Derived   IPut
	nodeId    string
	ctx       context.Context
	ctxCancel context.CancelFunc
}

func (this *Put) Open(root context.Context, nodeType int, putInterval int64) {
	this.ctx, this.ctxCancel = context.WithCancel(root)
	u, _ := uuid.NewV1()
	nodeId := fmt.Sprintf("%d-%s", nodeType, u.String())
	this.Derived.SetId(nodeId)
	xlog.Infoln("node id:", nodeId)
	go this.put(nodeType, putInterval)
}

func (this *Put) put(nodeType int, putInterval int64) {
	defer func() {
		if err := recover(); err != nil {
			xlog.Errorln("[except] ", err, "\n", string(debug.Stack()))
			if this.Derived.GetClient() != nil {
				go this.put(nodeType, putInterval)
			}
		}
	}()
	tick := time.NewTicker(time.Duration(putInterval) * time.Second)
	for {
		select {
		case <-tick.C:
			cli := this.Derived.GetClient()
			if cli == nil {
				return
			}
			resp, err := cli.Grant(context.TODO(), putInterval+5)
			if err != nil {
				xlog.Errorln(err)
			} else {
				var data string
				data, err = this.Derived.GetPutData()
				if err == nil {
					_, err = cli.Put(context.TODO(), this.Derived.Id(), data, clientv3.WithLease(resp.ID))
					if err != nil {
						xlog.Errorln(err)
					}
				}
			}
		case <-this.ctx.Done():
			return
		}
	}
}
