package internal

import (
	"github.com/fananchong/go-x/base"
	"github.com/fananchong/go-x/internal/k8s"
	"github.com/fananchong/multiconfig"
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
	path := base.ASSETS_PATH + "server_type.toml"
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
