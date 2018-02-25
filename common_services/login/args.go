package main

import (
	"github.com/fananchong/go-x/common"
)

type Args struct {
	common.ArgsBase
}

func (this *Args) OnInit() {
	this.Etcd.NodeType = 0                                                  // 不需要同步信息到discovery
	this.Etcd.WatchNodeTypes = []int{int(common.Gateway), int(common.Room)} // 监视服务节点类型
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
