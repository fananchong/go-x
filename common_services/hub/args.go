package main

import (
	"github.com/fananchong/go-x/common"
)

type Args struct {
	common.ArgsBase
	Hub ArgsHub
}

type ArgsHub struct {
}

func (this *Args) OnInit() {
}
