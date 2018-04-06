package discovery

import (
	"sync"

	godiscovery "github.com/fananchong/go-discovery"
	proto "github.com/golang/protobuf/proto"
)

type Node struct {
	godiscovery.Node
	info    ServerInfo
	Servers IServers
	mutex   sync.RWMutex
}

func (this *Node) InitPolicy(policy Policy) {
	switch policy {
	case Ordered:
		this.Servers = NewServersPolicyOrdered()
	case Random:
		this.Servers = NewServersPolicyRandom()
	case RoundRobin:
		this.Servers = NewServersPolicyRoundRobin()
	default:
	}
}

func (this *Node) GetBase() interface{} {
	return this
}

func (this *Node) OnNodeUpdate(nodeType int, id string, data []byte) {
	info := &ServerInfo{}
	err := proto.Unmarshal(data, info)
	if err == nil {
		this.Servers.Set(nodeType, id, info)
	} else {
		xlog.Errorln("[NODE] DATA ERROR!")
	}
}

func (this *Node) OnNodeJoin(nodeType int, id string, data []byte) {
	info := &ServerInfo{}
	err := proto.Unmarshal(data, info)
	if err == nil {
		this.Servers.Set(nodeType, id, info)
	} else {
		xlog.Errorln("[NODE] DATA ERROR!")
	}
}

func (this *Node) OnNodeLeave(nodeType int, id string) {
	this.Servers.Delete(nodeType, id)
}

func (this *Node) GetPutData() (string, error) {
	info := this.GetBaseInfo()
	data, err := proto.Marshal(&info)
	return string(data), err
}

// base info safe write / read

func (this *Node) GetBaseInfo() ServerInfo {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.info
}

func (this *Node) SetBaseInfoIP(externalIp string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.info.ExternalIp = externalIp
}

func (this *Node) SetBaseInfoType(t uint32) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.info.Type = t
}

func (this *Node) SetBaseInfoOrdered(ordered uint32) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.info.Ordered = ordered
}

var xnode *Node

func SetNode(node *Node) {
	xnode = node
}

func GetNode() *Node {
	return xnode
}

var xlog godiscovery.ILogger

func SetLogger(log godiscovery.ILogger) {
	xlog = log
	godiscovery.SetLogger(log)
}
