package common

import "flag"

type IArgs interface {
	Init()
	Parse()
	GetBase() *ArgsBase
}

type ArgsBase struct {
	EtcdHosts          string
	EtcdNodeType       int64
	EtcdWatchNodeTypes string
	EtcdPutInterval    int64
}

func (this *ArgsBase) Init() {
	// etcd
	flag.StringVar(&this.EtcdHosts, "etcdHosts", "192.168.1.4:12379,192.168.1.4:22379,192.168.1.4:32379", "etcd hosts")
	flag.Int64Var(&this.EtcdNodeType, "etcdNodeType", 1, "etcd node type")
	flag.StringVar(&this.EtcdWatchNodeTypes, "etcdWatchNodeTypes", "1,2,3,4", "etcd watch node type")
	flag.Int64Var(&this.EtcdPutInterval, "etcdPutInterval", 1, "etcd put interval")
}

func (this *ArgsBase) Parse() {

}

func (this *ArgsBase) GetBase() *ArgsBase {
	return this
}
