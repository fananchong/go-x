package main

import (
	"github.com/fananchong/go-x/common"
)

type Args struct {
	common.ArgsBase
	Base ArgsBase
}

type ArgsBase struct {
	ExternalIp string
}

func (this *Args) OnInit() {
}
