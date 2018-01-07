package common

type ArgsBase struct {
	Version    string   // 版本号
	ExternalIp string   // 外网ip（包括端口。格式 ip:port）
	IntranetIp string   // 内网ip（包括端口。格式 ip:port）
	Etcd       ArgsEtcd // etcd
}

type ArgsEtcd struct {
	Hosts          []string `default:""`  // etcd主机IP列表
	NodeType       int      `default:"0"` // 本节点类型。填0，则本节点不会上报自身信息给etcd。
	WatchNodeTypes []int    `default:""`  // 本节点要watch其他节点的节点类型
	PutInterval    int      `default:"1"` // 本节点上报信息间隔，单位秒
}
