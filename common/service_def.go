package common

import (
	"github.com/fananchong/go-x/common/k8s"
	"github.com/fananchong/multiconfig"
)

type CommonServerType int

const (
	Client            CommonServerType = iota // 0
	Login                                     // 1
	Gateway                                   // 2
	Mgr                                       // 3
	COMMON_SERVER_END = 9                     // 9
)

type ServerTypeConfig struct {
	ServerType []ServerTypeInfo
}

type ServerTypeInfo struct {
	Type      int
	Name      string
	Namespace string
}

func initServerType() {
	path := GetAssetsPath() + "server_type.toml"
	cfg := &ServerTypeConfig{}
	m := &multiconfig.TOMLLoader{
		Path: path,
	}
	err := m.Load(cfg)
	if err != nil {
		panic(err)
	}
	for _, v := range cfg.ServerType {
		k8s.RegisterNodeType(v.Type, v.Namespace, v.Name)
	}
}
