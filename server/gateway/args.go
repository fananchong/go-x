package main

import (
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/server/def"
)

var (
	xargs *Args = NewArgs()
)

type Args struct {
	common.ArgsBase
}

func NewArgs() *Args {
	return &Args{}
}

func (this *Args) OnInit() {
	this.Etcd.NodeType = int(def.Gateway)
	this.Etcd.WatchNodeTypes = []int{int(def.Base)} // 监视服务节点类型
}

func (this *Args) GetDerived() common.IArgs {
	return this
}
