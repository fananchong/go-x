package main

import (
	"def"

	"github.com/fananchong/go-x/common/discovery"
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
