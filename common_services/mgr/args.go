package main

import (
	"github.com/fananchong/go-x/base"
)

type Args struct {
	base.ArgsBase
	Mgr ArgsMgr
}

type ArgsMgr struct {
}

func (this *Args) OnInit() {
}
