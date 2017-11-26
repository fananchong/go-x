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
	Derived  IPut
	nodeId   string
	tick     *time.Ticker
	chanStop chan int
}

func (this *Put) Open(nodeType int, putInterval int64) {
	nodeId := fmt.Sprintf("%d-%s", nodeType, uuid.NewV1().String())
	this.Derived.SetId(nodeId)
	this.Derived.GetLogger().Infoln("node id:", nodeId)
	go this.put(nodeType, putInterval)
}

func (this *Put) put(nodeType int, putInterval int64) {
	defer func() {
		if err := recover(); err != nil {
			this.Derived.GetLogger().Errorln("[异常] ", err, "\n", string(debug.Stack()))
		}
		this.Derived.Close()
	}()
	this.tick = time.NewTicker(time.Duration(putInterval) * time.Second)
	for {
		select {
		case <-this.tick.C:
			cli := this.Derived.GetClient()
			if cli == nil {
				return
			}
			resp, err := cli.Grant(context.TODO(), putInterval+5)
			if err != nil {
				this.Derived.GetLogger().Errorln(err)
			} else {
				var data string
				data, err = this.Derived.GetPutData()
				if err == nil {
					_, err = cli.Put(context.TODO(), this.Derived.Id(), data, clientv3.WithLease(resp.ID))
					if err != nil {
						this.Derived.GetLogger().Errorln(err)
					}
				}
			}
		case <-this.chanStop:
			return
		}
	}
}

func (this *Put) Close() {
	if this.tick != nil {
		this.tick.Stop()
		this.tick = nil
	}
	this.chanStop <- 1
}
