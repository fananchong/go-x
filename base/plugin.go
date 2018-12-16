package base

type Plugin interface {
	Init() bool
	Start() bool
	Close()
}
