package internal

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fananchong/go-x/base"
	"github.com/fananchong/multiconfig"
)

var (
	_ = flag.String("assets", "", "path of assets")
	_ = flag.String("app", "", "app name")
)

func parseArgs(derived interface{}) {
	for i, v := range os.Args {
		if v == "-assets" || v == "--assets" {
			base.ASSETS_PATH = os.Args[i+1] + "/"
		}
		if v == "-h" || v == "-help" || v == "--help" || v == "/?" {
			fl := &multiconfig.FlagLoader{}
			fl.Load(derived)
			flag.CommandLine.PrintDefaults()
			os.Exit(0)
		}
	}
	if base.ASSETS_PATH == "" {
		base.ASSETS_PATH = "./"
	}
	dir, err := filepath.Abs(filepath.Dir(base.ASSETS_PATH))
	if err != nil {
		fmt.Println("no find assets path, path: " + base.ASSETS_PATH)
	}
	fmt.Println("Assets Path:", dir)
	base.ASSETS_PATH = dir + "/"
	cfg := base.ASSETS_PATH + "config.toml"
	_, err = os.Stat(cfg)
	if !(err == nil || os.IsExist(err)) {
		fmt.Println("no find config.toml " + cfg)
	}
	m := multiconfig.NewWithPath(cfg)
	m.MustLoad(derived)
}

func GetAppName() string {
	if base.APP_NAME == "" {
		for i, v := range os.Args {
			if v == "-app" || v == "--app" {
				base.APP_NAME = os.Args[i+1]
				break
			}
		}
	}
	return base.APP_NAME
}
