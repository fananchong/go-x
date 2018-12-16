package base

type CommonServerType int

const (
	Client            CommonServerType = iota // 0
	Login                                     // 1
	Gateway                                   // 2
	Mgr                                       // 3
	COMMON_SERVER_END = 9                     // 9
)
