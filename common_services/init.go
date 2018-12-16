package service

import "github.com/fananchong/go-x/base"

func init() {
	LoadIpConfig(base.ASSETS_PATH + "ip.toml")
}
