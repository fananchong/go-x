package main

import (
	"def"

	"github.com/fananchong/go-x/common"
)

type Args struct {
	common.ArgsBase
}

func (this *Args) OnInit() {
	this.Etcd.NodeType = int(def.Gateway)
	this.Etcd.WatchNodeTypes = []int{int(def.Base)} // 监视服务节点类型
}

var (
	xargs *Args = NewArgs()
)

func NewArgs() *Args {
	return &Args{}
}

func (this *Args) GetDerived() common.IArgs {
	return this
}
