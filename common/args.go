package common

import (
	"github.com/fananchong/multiconfig"
)

type ArgsBase struct {
	CfgPath    string   // 配置路径
	Version    string   // 版本号
	ExternalIp string   // 外网ip（包括端口。格式 ip:port）
	IntranetIp string   // 内网ip（包括端口。格式 ip:port）
	Etcd       ArgsEtcd // etcd
}

type ArgsEtcd struct {
	Hosts          []string `default:[]`
	NodeType       int      `default:"0"`
	WatchNodeTypes []int    `default:[]`
	PutInterval    int      `default:"1"`
}

type IArgs interface {
	IArgsBase
	OnInit()
	GetDerived() IArgs
}

type IArgsBase interface {
	GetBase() *ArgsBase
	Init(derived IArgs)
}

func (this *ArgsBase) Init(derived IArgs) {
	m := multiconfig.NewWithPath("config.toml")
	m.MustLoad(derived)
	derived.OnInit()
}

func (this *ArgsBase) GetBase() *ArgsBase {
	return this
}
