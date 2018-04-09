package common

type ServerType int

const (
	Client  ServerType = iota // 0
	Login                     // 1
	Gateway                   // 2
	Lobby                     // 3
	Room                      // 4
	Hub                       // 5
)
