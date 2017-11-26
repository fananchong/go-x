package main

import (
	"github.com/fananchong/go-x/common/discovery"
	"github.com/fananchong/go-x/def"
)

var (
	xnode *Node = NewNode()
)

type Node struct {
	discovery.Node
}

func NewNode() *Node {
	node := &Node{}
	node.Info.Type = uint32(def.Login)
	node.InitPolicy(discovery.Ordered)
	return node
}
