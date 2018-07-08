package service

import "github.com/fananchong/go-x/common"

func init() {
	LoadIpConfig(common.GetAssetsPath() + "ip.toml")
}

