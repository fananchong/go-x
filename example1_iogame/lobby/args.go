package main

import (
	"github.com/fananchong/go-x/common"
)

var DbAccount string

type Args struct {
	common.ArgsBase
	Lobby ArgsLobby
}

type ArgsLobby struct {
}

func (this *Args) OnInit() {
	DbAccount = this.ArgsBase.DbAccount.Name
}
