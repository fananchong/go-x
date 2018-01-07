package main

import (
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/server/def"
)

type Args struct {
	common.ArgsBase
}

func (this *Args) OnInit() {
	this.Etcd.NodeType = int(def.Room)
}

var (
	xargs *Args = NewArgs()
)

func NewArgs() *Args {
	return &Args{}
}

func (this *Args) GetDerived() common.IArgs {
	return this
}
