package common

type ArgsBase struct {
	Common    ArgsCommon  // 一些基础参数
	Pending   ArgsPending // 悬而未决的配置，OnInit()函数内用户最终设置
	Etcd      ArgsEtcd    // Etcd配置
	DbAccount ArgsRedis   // 帐号数据库（Redis）
	DbToken   ArgsRedis   // Token数据库（Redis）
	Login     ArgsLogin   // 登录服务配置
}

type ArgsCommon struct {
	Version  string `default:""` // 版本号
	LogDir   string `default:""` // log路径
	LogLevel int    `default:2`  // log等级
}

type ArgsPending struct {
	ExternalIp string
	IntranetIp string
}

type ArgsEtcd struct {
	Hosts          []string `default:""`  // etcd主机IP列表
	NodeType       int      `default:"0"` // 本节点类型。填0，则本节点不会上报自身信息给etcd。
	WatchNodeTypes []int    `default:""`  // 本节点要watch其他节点的节点类型
	PutInterval    int      `default:"1"` // 本节点上报信息间隔，单位秒
}

type ArgsRedis struct {
	Name     string   `default:""`
	Addrs    []string `default:""`
	Password string   `default:""`
	DBIndex  int      `default:0`
}

type ArgsLogin struct {
	Listen string `default:":8000"`
	Sign1  string `default:""`
	Sign2  string `default:""`
	Sign3  string `default:""`
}
