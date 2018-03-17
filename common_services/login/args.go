package main

import (
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/discovery"
)

type Args struct {
	common.ArgsBase
	Login ArgsLogin // 登录服务配置
}

type ArgsLogin struct {
	Listen string `default:":8000"`
	Sign1  string `default:""`
	Sign2  string `default:""`
	Sign3  string `default:""`
}

func (this *Args) OnInit() {
	this.Pending.NodeType = 0                                // 不需要同步信息到discovery
	this.Pending.WatchNodeTypes = []int{int(common.Gateway)} // 监视服务节点类型
	xnode.SetBaseInfoType(uint32(common.Login))              // 设置本节点类型
	xnode.InitPolicy(discovery.Ordered)                      // 设置选取服务节点策略
}

func NewArgs() *Args {
	return &Args{}
}

func (this *Args) GetDerived() common.IArgs {
	return this
}
