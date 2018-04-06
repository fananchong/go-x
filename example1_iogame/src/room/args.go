package main

import (
	"github.com/fananchong/go-x/common"
)

type Args struct {
	common.ArgsBase
}

func (this *Args) OnInit() {
}

func NewArgs() *Args {
	return &Args{}
}

func (this *Args) GetDerived() common.IArgs {
	return this
}
