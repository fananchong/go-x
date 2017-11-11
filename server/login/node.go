package main

import godiscovery "github.com/fananchong/go-discovery"

var (
	xnode *Node = NewNode()
)

type Node struct {
	godiscovery.Node
}

func NewNode() *Node {
	return &Node{}
}

func (this *Node) OnNodeUpdate(nodeType int, id string, data []byte) {
}

func (this *Node) OnNodeJoin(nodeType int, id string, data []byte) {
}

func (this *Node) OnNodeLeave(nodeType int, id string) {
}

func (this *Node) GetPutData() string {
	return ""
}
