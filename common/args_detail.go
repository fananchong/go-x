package common

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fananchong/multiconfig"
)

type IArgs interface {
	IArgsBase
	OnInit()
}

type IArgsBase interface {
	GetBase() *ArgsBase
	Init(derived IArgs)
}

var (
	assetsPath = flag.String("assets", "", "path of assets")
)

func (this *ArgsBase) Init(derived IArgs) {
	index := 0
	for i, v := range os.Args {
		if v == "-assets" || v == "--assets" {
			index = i + 1
		}
	}
	assetsPath := ""
	if index != 0 && index < len(os.Args) {
		assetsPath = os.Args[index] + "/"
	}
	if assetsPath == "" {
		assetsPath = "./"
	}
	dir, err := filepath.Abs(filepath.Dir(assetsPath))
	if err != nil {
		panic("no find assets path, path: " + assetsPath)
	}
	fmt.Println("Assets Path:", dir)
	cfg := assetsPath + "config.toml"
	_, err = os.Stat(cfg)
	if !(err == nil || os.IsExist(err)) {
		panic("no find config.toml, path: " + cfg)
	}
	m := multiconfig.NewWithPath(cfg)
	m.MustLoad(derived)
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

func GetAssetsPath() string {
	return *assetsPath
}
