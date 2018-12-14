package main

import (
	"plugin"

	"github.com/fananchong/go-x/common"
)

func main() {
	app := common.GetAppName()
	p, err := plugin.Open(app + ".so")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("Run")
	if err != nil {
		panic(err)
	}
	f.(func())()
}
