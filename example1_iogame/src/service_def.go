package iogame

// 公共服务类型，请参见：go-x/common/service_def.go

type ServerType int

const (
	Client ServerType = iota // 0
	Lobby             = 10   // 10
	Room              = 11   // 11
)
