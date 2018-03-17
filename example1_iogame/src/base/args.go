package main

import (
	"github.com/fananchong/go-x/common"
)

type Args struct {
	common.ArgsBase
}

func (this *Args) OnInit() {
	this.Pending.NodeType = int(common.Base)   // 设置本节点类型。不为0，则上报自身节点信息到discovery。
	xnode.SetBaseInfoType(uint32(common.Base)) // 设置本节点类型
}

func NewArgs() *Args {
	return &Args{}
}

func (this *Args) GetDerived() common.IArgs {
	return this
}
