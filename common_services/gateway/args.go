package main

import (
	"github.com/fananchong/go-x/common"
)

type Args struct {
	common.ArgsBase
	Gateway ArgsGateway
}

type ArgsGateway struct {
	Connect []int
}

func (this *Args) OnInit() {

}
