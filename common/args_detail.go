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
	_          = flag.String("assets", "", "path of assets")
	_          = flag.String("app", "", "app name")
	assetsPath string
	app        string
)

func (this *ArgsBase) Init(derived IArgs) {
	for i, v := range os.Args {
		if v == "-assets" || v == "--assets" {
			assetsPath = os.Args[i+1] + "/"
		}
		if v == "-h" || v == "-help" || v == "--help" || v == "/?" {
			fl := &multiconfig.FlagLoader{}
			fl.Load(derived)
			flag.CommandLine.PrintDefaults()
			os.Exit(0)
		}
	}
	if assetsPath == "" {
		assetsPath = "./"
	}
	dir, err := filepath.Abs(filepath.Dir(assetsPath))
	if err != nil {
		fmt.Println("no find assets path, path: " + assetsPath)
	}
	fmt.Println("Assets Path:", dir)
	assetsPath = dir + "/"
	cfg := assetsPath + "config.toml"
	_, err = os.Stat(cfg)
	if !(err == nil || os.IsExist(err)) {
		fmt.Println("no find config.toml " + cfg)
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
	return assetsPath
}

func GetAppName() string {
	if app == "" {
		for i, v := range os.Args {
			if v == "-app" || v == "--app" {
				app = os.Args[i+1]
				break
			}
		}
	}
	return app
}
