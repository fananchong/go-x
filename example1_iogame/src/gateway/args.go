package main

import (
	"github.com/fananchong/go-x/common"
)

type Args struct {
	common.ArgsBase
}

func (this *Args) OnInit() {
	this.Etcd.NodeType = int(common.Gateway)
	this.Etcd.WatchNodeTypes = []int{int(common.Base)} // 监视服务节点类型
}

func NewArgs() *Args {
	return &Args{}
}

func (this *Args) GetDerived() common.IArgs {
	return this
}
