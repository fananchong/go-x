package discovery

import (
	"github.com/fananchong/go-discovery"
	proto "github.com/golang/protobuf/proto"
)

type Node struct {
	godiscovery.Node
	Info ServerInfo
}

func (this *Node) GetBase() interface{} {
	return this
}

func (this *Node) OnNodeUpdate(nodeType int, id string, data []byte) {
}

func (this *Node) OnNodeJoin(nodeType int, id string, data []byte) {
}

func (this *Node) OnNodeLeave(nodeType int, id string) {
}

func (this *Node) GetPutData() (string, error) {
	info := this.Info
	data, err := proto.Marshal(&info)
	return string(data), err
}
