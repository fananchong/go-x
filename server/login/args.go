package main

import (
	"flag"
	"fmt"

	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/def"
)

var (
	xargs *Args = NewArgs()
)

type Args struct {
	common.ArgsBase
	Listen string
}

func NewArgs() *Args {
	return &Args{}
}

func (this *Args) Init() {
	this.ArgsBase.Init()
	flag.StringVar(&this.Listen, "listen", ":8000", "listen address")
}

func (this *Args) Parse() {
	this.ArgsBase.Parse()
	this.EtcdNodeType = 0                                                 // 不需要同步信息到discovery
	this.EtcdWatchNodeTypes = fmt.Sprintf("%v,%v", def.Gateway, def.Room) // 监视服务节点类型
}
