package main

import "github.com/fananchong/go-x/base"

var externArgs ArgsLogin

type Args struct {
	base.ArgsBase
	Login ArgsLogin // 登录服务配置
}

type ArgsLogin struct {
	Listen string `default:":8000"`
	Sign1  string `default:""`
	Sign2  string `default:""`
	Sign3  string `default:""`
}

func (this *Args) OnInit() {
	this.Pending.WatchNodeTypes = append(this.Pending.WatchNodeTypes, int(base.Gateway)) // 监视服务节点类型
	externArgs = this.Login
}
