package main

import "flag"

type ArgsRaw struct {
	EtcdHosts      string
	NodeType       int64
	WatchNodeTypes string
	PutInterval    int64
}

type Args struct {
	ArgsRaw
}

func NewArgs() *Args {
	return &Args{}
}

func (this *Args) Init() {

	// etcd
	flag.StringVar(&this.EtcdHosts, "etcd_hosts", "192.168.1.4:12379,192.168.1.4:22379,192.168.1.4:32379", "etcd hosts")
	flag.Int64Var(&this.NodeType, "nodeType", 1, "node type")
	flag.StringVar(&this.WatchNodeTypes, "watchNodeTypes", "1,2,3,4", "watch node type")
	flag.Int64Var(&this.PutInterval, "putInterval", 1, "put interval")
}

var (
	xargs *Args = NewArgs()
)
