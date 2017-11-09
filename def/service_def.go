package def

type ServerType int

const (
	_       ServerType = iota // 0
	Base                      // 1
	Gateway                   // 2
	Login                     // 3
	Room                      // 4
)
