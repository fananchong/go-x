package common

type ArgsBase struct {
	Common    ArgsCommon  // 一些基础参数
	Pending   ArgsPending // 悬而未决的配置，OnInit()函数内用户最终设置。不同的节点会有针对性的设置。
	Etcd      ArgsEtcd    // Etcd配置
	DbAccount ArgsRedis   // 帐号数据库（Redis）
	DbToken   ArgsRedis   // Token数据库（Redis）
	DbServer  ArgsRedis   // Server数据库（Redis）
}

type ArgsCommon struct {
	Version       string `default:""` // 版本号
	LogDir        string `default:""` // log路径
	LogLevel      int    `default:2`  // log等级
	Debug         bool   `default:0`  // debug版本标志
	IntranetToken string `default:""` // 内部服务器验证TOKEN
}

type ArgsPending struct {
	NodeType       int   `default:"0"` // 本节点类型。填0，则本节点不会上报自身信息给etcd。
	WatchNodeTypes []int `default:""`  // 本节点要watch其他节点的节点类型
}

type ArgsEtcd struct {
	Hosts       []string `default:""`  // etcd主机IP列表
	PutInterval int      `default:"1"` // 本节点上报信息间隔，单位秒
	WhatsMyIP   string   `default:""`  // whatsmyip主机IP
}

type ArgsRedis struct {
	Name     string   `default:""`
	Addrs    []string `default:""`
	Password string   `default:""`
	DBIndex  int      `default:0`
}
