package godiscovery

import (
	"context"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type INode interface {
	Id() string
	SetId(id string)
	Init(inst interface{})
	Open(hosts []string, nodeType int, watchNodeTypes []int, putInterval int64)
	Close()
	GetClient() *clientv3.Client
	GetBase() interface{}
}

type Node struct {
	Watch
	Put
	client         *clientv3.Client
	hosts          []string
	nodeType       int
	watchNodeTypes []int
	putInterval    int64
	mutex          sync.RWMutex
	ctx            context.Context
	ctxCancel      context.CancelFunc
}

func (this *Node) Init(inst interface{}) {
	this.ctx, this.ctxCancel = context.WithCancel(context.Background())
	this.Watch.Derived = inst.(IWatch)
	this.Put.Derived = inst.(IPut)
}

func (this *Node) Open(hosts []string, nodeType int, watchNodeTypes []int, putInterval int64) {
	this.hosts = hosts
	this.nodeType = nodeType
	this.watchNodeTypes = watchNodeTypes
	this.putInterval = putInterval
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   hosts,
		DialTimeout: 5 * time.Second,
	})
	this.client = cli
	if err != nil {
		xlog.Errorln(err)
		if cli != nil {
			cli.Close()
		}
		go this.reopen()
		return
	}
	if len(watchNodeTypes) != 0 {
		this.Watch.Open(this.ctx, watchNodeTypes)
	}
	if nodeType != 0 {
		this.Put.Open(this.ctx, nodeType, putInterval)
	}
}

func (this *Node) OpenByStr(hostsStr string, nodeType int, watchNodeTypesStr string, putInterval int64) {
	hosts := strings.Split(hostsStr, ",")
	var watchNodeTypes []int = make([]int, 0)
	if watchNodeTypesStr != "" {
		for _, val := range strings.Split(watchNodeTypesStr, ",") {
			v, _ := strconv.Atoi(val)
			if v == 0 {
				continue
			}
			watchNodeTypes = append(watchNodeTypes, v)
		}
	}
	this.Open(hosts, nodeType, watchNodeTypes, putInterval)
}

func (this *Node) Close() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.ctxCancel != nil {
		this.ctxCancel()
		this.ctxCancel = nil
	}
	if this.client != nil {
		this.client.Close()
		this.client = nil
	}
}

func (this *Node) reopen() {
	xlog.Infoln("reopen after 5 sec.")
	t := time.NewTimer(5 * time.Second)
	select {
	case <-t.C:
		this.Open(this.hosts, this.nodeType, this.watchNodeTypes, this.putInterval)
	}
}

func (this *Node) Id() string {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.Put.nodeId
}

func (this *Node) SetId(id string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.Put.nodeId = id
}

func (this *Node) GetClient() *clientv3.Client {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.client
}

// 子类可以根据需要重载下面的方法
//     注意 OnNodeUpdate、OnNodeJoin、OnNodeLeave、GetPutData 在内部协程被调用，请注意多协程安全！！！
func (this *Node) OnNodeUpdate(nodeType int, id string, data []byte) {

}

func (this *Node) OnNodeJoin(nodeType int, id string, data []byte) {

}

func (this *Node) OnNodeLeave(nodeType int, id string) {

}

func (this *Node) GetPutData() (string, error) {
	return "", nil
}

func (this *Node) GetBase() interface{} {
	return this
}
