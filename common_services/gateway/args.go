package main

import "github.com/fananchong/go-x/base"

type Args struct {
	base.ArgsBase
	Gateway ArgsGateway
}

type ArgsGateway struct {
	Connect []int
}

func (this *Args) OnInit() {

}
