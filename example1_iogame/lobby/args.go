package main

import (
	"github.com/fananchong/go-x/common"
)

type Args struct {
	common.ArgsBase
	Lobby ArgsLobby
}

type ArgsLobby struct {
}

func (this *Args) OnInit() {
}
