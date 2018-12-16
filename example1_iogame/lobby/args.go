package main

import (
	"github.com/fananchong/go-x/base"
)

var DbAccount string

type Args struct {
	base.ArgsBase
	Lobby ArgsLobby
}

type ArgsLobby struct {
}

func (this *Args) OnInit() {
	DbAccount = this.ArgsBase.DbAccount.Name
}
