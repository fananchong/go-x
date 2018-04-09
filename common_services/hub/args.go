package main

import (
	"github.com/fananchong/go-x/common"
)

type Args struct {
	common.ArgsBase
	Hub ArgsHub
}

type ArgsHub struct {
	ExternalIp string
}

func (this *Args) OnInit() {
}
