package godiscovery

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	uuid "github.com/satori/go.uuid"
)

type INode interface {
	Id() string
	SetId(id string)
	Init(inst interface{})
	Open(hosts []string, whatsmyip string, nodeType int, watchNodeTypes []int, putInterval int64)
	Close()
	GetClient() *clientv3.Client
	GetBase() interface{}
}

type Node struct {
	Watch
	Put
	Port
	client         *clientv3.Client
	hosts          []string
	whatsmyipHost  string
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

func (this *Node) Open(hosts []string, whatsmyipHost string, nodeType int, watchNodeTypes []int, putInterval int64) {
	this.hosts = hosts
	this.whatsmyipHost = whatsmyipHost
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

	if err := this.Port.Init(this.ctx, this.client); err != nil {
		xlog.Errorln(err)
		if cli != nil {
			cli.Close()
		}
		go this.reopen()
		return
	}

	if nodeType != 0 {
		if whatsmyipHost != "" {
			resp, err := http.Get("http://" + whatsmyipHost)
			if err != nil {
				xlog.Errorln(err)
				if cli != nil {
					cli.Close()
				}
				go this.reopen()
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				xlog.Errorln(err)
				if cli != nil {
					cli.Close()
				}
				go this.reopen()
				return
			}
			this.Put.nodeIP = fmt.Sprintf("%s:%d", string(body), this.GetPort())
		}
		if err := this.Put.Open(this.ctx, nodeType, putInterval); err != nil {
			xlog.Errorln(err)
			if cli != nil {
				cli.Close()
			}
			go this.reopen()
			return
		}
	}
	if len(watchNodeTypes) != 0 {
		this.Watch.Open(this.ctx, watchNodeTypes)
	}
}

func (this *Node) OpenByStr(hostsStr string, whatsmyipHost string, nodeType int, watchNodeTypesStr string, putInterval int64) {
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
	this.Open(hosts, whatsmyipHost, nodeType, watchNodeTypes, putInterval)
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
		this.Open(this.hosts, this.whatsmyipHost, this.nodeType, this.watchNodeTypes, this.putInterval)
	}
}

func (this *Node) Id() string {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.Put.nodeId
}

func (this *Node) Ip() string {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.Put.nodeIP
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
func (this *Node) OnNodeUpdate(nodeIP string, nodeType int, id string, data []byte) {

}

func (this *Node) OnNodeJoin(nodeIP string, nodeType int, id string, data []byte) {

}

func (this *Node) OnNodeLeave(nodeType int, id string) {

}

func (this *Node) GetPutData() (string, error) {
	return "", nil
}

func (this *Node) NewNodeId() (string, error) {
	id, err := uuid.NewV1()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (this *Node) GetBase() interface{} {
	return this
}
