package gotcp

import (
	"github.com/golang/protobuf/proto"
)

func (this *Session) SendMsg(cmd uint64, msg proto.Message) bool {
	data, flag, err := EncodeCmd(cmd, msg)
	if err != nil {
		return false
	}
	return this.Send(data, flag)
}
