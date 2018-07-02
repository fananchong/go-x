package common

import "github.com/fananchong/go-x/common/k8s"

type ServerType int

const (
	Client  ServerType = iota // 0
	Login                     // 1
	Gateway                   // 2
	Lobby                     // 3
	Room                      // 4
	Hub                       // 5
)

func init() {
	const ns = "go-x"
	k8s.RegisterNodeType(int(Login), ns, "login")
	k8s.RegisterNodeType(int(Gateway), ns, "gateway")
	k8s.RegisterNodeType(int(Lobby), ns, "lobby")
	k8s.RegisterNodeType(int(Room), ns, "room")
	k8s.RegisterNodeType(int(Hub), ns, "hub")
}
