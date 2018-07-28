package common

type MessagesFunc func(data []byte, flag byte)

type Messages struct {
	Handlers map[uint32]MessagesFunc
	Cmds     []uint32
}

func NewMessages() *Messages {
	return &Messages{
		Handlers: make(map[uint32]MessagesFunc),
	}
}

func (this *Messages) RegisterMessage(cmd uint32, f MessagesFunc) {
	this.Handlers[cmd] = f
	this.Cmds = append(this.Cmds, cmd)
}
