package common

import "flag"

type IArgs interface {
	Init()
	Parse()
	GetBase() *ArgsBase
}

type ArgsBase struct {
	// version
	Version string // 版本号

	// ip
	ExternalIp string // 外网ip（包括端口。格式 ip:port）
	IntranetIp string // 内网ip（包括端口。格式 ip:port）

	// etcd
	EtcdHosts          string
	EtcdNodeType       int64
	EtcdWatchNodeTypes string
	EtcdPutInterval    int64
}

func (this *ArgsBase) Init() {
	// version
	flag.StringVar(&this.Version, "version", "", "version")

	// ip
	flag.StringVar(&this.ExternalIp, "externalIp", "", "external ip")
	flag.StringVar(&this.IntranetIp, "intranetIp", "", "intranet ip")

	// etcd
	flag.StringVar(&this.EtcdHosts, "etcdHosts", "192.168.1.4:12379,192.168.1.4:22379,192.168.1.4:32379", "etcd hosts")
	flag.Int64Var(&this.EtcdNodeType, "etcdNodeType", 1, "etcd node type")
	flag.StringVar(&this.EtcdWatchNodeTypes, "etcdWatchNodeTypes", "", "etcd watch node type")
	flag.Int64Var(&this.EtcdPutInterval, "etcdPutInterval", 1, "etcd put interval")
}

func (this *ArgsBase) Parse() {

}

func (this *ArgsBase) GetBase() *ArgsBase {
	return this
}
