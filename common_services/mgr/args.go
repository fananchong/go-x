package main

import (
	"github.com/fananchong/go-x/common"
)

type Args struct {
	common.ArgsBase
	Mgr ArgsMgr
}

type ArgsMgr struct {
}

func (this *Args) OnInit() {
}
