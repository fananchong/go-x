package main

import (
	"plugin"

	"github.com/fananchong/go-x/base"
	"github.com/fananchong/go-x/internal"
)

func main() {
	appName := internal.GetAppName()
	p, err := plugin.Open(appName + ".so")
	if err != nil {
		panic(err)
	}
	obj, err := p.Lookup("MainObj")
	if err != nil {
		panic(err)
	}

	app := NewApp(obj.(base.Plugin))
	//app.Run()
	app.Close()
}
