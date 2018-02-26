package main

import (
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/discovery"
)

type Node struct {
	discovery.Node
}

func NewNode() *Node {
	node := &Node{}
	node.SetBaseInfoType(uint32(common.Login))
	node.InitPolicy(discovery.Ordered)
	return node
}
