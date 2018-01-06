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
	node.SetBaseInfoType(uint32(def.Room))
	return node
}
