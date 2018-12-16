package discovery

import (
	"sync"
	"time"

	"github.com/fananchong/go-x/internal/k8s"
)

type Node struct {
	k8s.Node
	Servers IServers
	mutex   sync.RWMutex
}

func NewNode() *Node {
	this := &Node{}
	this.Servers = NewServersPolicyRoundRobin()
	return this
}

func (this *Node) Init(nodeType int, watchNodeTypes []int, d time.Duration, inst k8s.INode) error {
	SetNode(this)
	return this.Node.Init(nodeType, watchNodeTypes, d, inst)
}

func (this *Node) OnNodeJoin(endpoint *k8s.Endpoint) {
	this.Servers.Set(endpoint.NodeType, endpoint.Id(), endpoint)
}

func (this *Node) OnNodeLeave(endpoint *k8s.Endpoint) {
	this.Servers.Delete(endpoint.NodeType, endpoint.Id())
	this.Node.OnNodeLeave(endpoint)
}

/// ==================================================

func (this *Node) GetBase() interface{} {
	return this
}

var xnode *Node

func SetNode(node *Node) {
	xnode = node
}

func GetNode() *Node {
	return xnode
}
