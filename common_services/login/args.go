package main

import (
	"github.com/fananchong/go-x/common"
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
	this.Pending.NodeType = 0                                                              // 不需要同步信息到discovery
	this.Pending.WatchNodeTypes = append(this.Pending.WatchNodeTypes, int(common.Gateway)) // 监视服务节点类型
}
