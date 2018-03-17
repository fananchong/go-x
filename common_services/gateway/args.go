package main

import (
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/discovery"
)

type Args struct {
	common.ArgsBase
	Gateway ArgsGateway
}

type ArgsGateway struct {
	ExternalIp string
	IntranetIp string
	Connect    []int
}

func (this *Args) OnInit() {
	this.Pending.NodeType = int(common.Gateway)                // 设置本节点类型。不为0，则上报自身节点信息到discovery。
	this.Pending.WatchNodeTypes = append(this.Gateway.Connect) // 监视服务节点类型
	this.Pending.ExternalIp = this.Gateway.ExternalIp          // 对外地址
	this.Pending.IntranetIp = this.Gateway.IntranetIp          // 对内地址
	xnode.SetBaseInfoType(uint32(common.Gateway))              // 设置本节点类型
	xnode.InitPolicy(discovery.Random)                         // 设置选取服务节点策略
}

func NewArgs() *Args {
	return &Args{}
}

func (this *Args) GetDerived() common.IArgs {
	return this
}
