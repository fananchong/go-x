package godiscovery

import (
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
	OpenByStr(hostsStr string, nodeType int, watchNodeTypesStr string, putInterval int64)
	Close()
	GetClient() *clientv3.Client
	GetLogger() ILogger
	SetLogger(log ILogger)
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
	mutexClose     sync.Mutex
	mutexVar       sync.RWMutex
	log            ILogger
}

func (this *Node) Init(inst interface{}) {
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
		this.log.Errorln(err)
		go this.reopen()
		return
	}
	if len(watchNodeTypes) != 0 {
		this.Watch.Open(watchNodeTypes)
	}
	if nodeType != 0 {
		this.Put.Open(nodeType, putInterval)
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
	this.mutexClose.Lock()
	defer this.mutexClose.Unlock()
	if this.client != nil {
		this.client.Close()
		this.client = nil
	}
	this.Put.Close()
	this.Watch.Close()
}

func (this *Node) reopen() {
	this.log.Infoln("reopen after 5 sec.")
	t := time.NewTimer(5 * time.Second)
	select {
	case <-t.C:
		this.Open(this.hosts, this.nodeType, this.watchNodeTypes, this.putInterval)
	}
}

func (this *Node) Id() string {
	this.mutexVar.RLock()
	defer this.mutexVar.RUnlock()
	return this.Put.nodeId
}

func (this *Node) SetId(id string) {
	this.mutexVar.Lock()
	defer this.mutexVar.Unlock()
	this.Put.nodeId = id
}

func (this *Node) GetClient() *clientv3.Client {
	this.mutexVar.RLock()
	defer this.mutexVar.RUnlock()
	return this.client
}

func (this *Node) GetLogger() ILogger {
	this.mutexVar.RLock()
	defer this.mutexVar.RUnlock()
	return this.log
}

func (this *Node) SetLogger(log ILogger) {
	this.mutexVar.Lock()
	defer this.mutexVar.Unlock()
	this.log = log
}

// 子类可以根据需要重载下面的方法
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
