package common

import (
	"flag"
	"fmt"
	"os"

	"github.com/fananchong/multiconfig"
)

type IArgs interface {
	IArgsBase
	OnInit()
	GetDerived() IArgs
}

type IArgsBase interface {
	GetBase() *ArgsBase
	Init(derived IArgs)
}

var (
	cfgpath = flag.String("cfg", "", "path of cfg file")
)

func (this *ArgsBase) Init(derived IArgs) {
	index := 0
	for i, v := range os.Args {
		if v == "-cfg" || v == "--cfg" {
			index = i + 1
		}
	}
	cfg := ""
	if index != 0 && index < len(os.Args) {
		cfg = os.Args[index]
		fmt.Println("cfg file:", cfg)
	}
	m := multiconfig.NewWithPath(cfg)
	m.MustLoad(derived)
	derived.OnInit()
}

func (this *ArgsBase) GetBase() *ArgsBase {
	return this
}

var xargs *ArgsBase

func SetArgs(args *ArgsBase) {
	xargs = args
}

func GetArgs() *ArgsBase {
	return xargs
}
