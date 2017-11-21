package main

import (
	"fmt"

	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/def"
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

func (this *Args) Init() {
	this.ArgsBase.Init()
}

func (this *Args) Parse() {
	this.ArgsBase.Parse()
	this.EtcdNodeType = int64(def.Gateway)
	this.EtcdWatchNodeTypes = fmt.Sprintf("%v", def.Base) // 监视服务节点类型
}
