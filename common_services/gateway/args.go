package main

import (
	"github.com/fananchong/go-x/common"
)

type Args struct {
	common.ArgsBase
	Gateway ArgsGateway
}

type ArgsGateway struct {
	ExternalIp string
	IntranetIp string
	Connect    []int
}

func (this *Args) OnInit() {

}
