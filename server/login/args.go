package main

import (
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/server/def"
)

type Args struct {
	common.ArgsBase
	Listen string `default:":8000"`
}

func (this *Args) OnInit() {
	this.Etcd.NodeType = 0                                            // 不需要同步信息到discovery
	this.Etcd.WatchNodeTypes = []int{int(def.Gateway), int(def.Room)} // 监视服务节点类型
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
