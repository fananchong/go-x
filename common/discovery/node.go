package discovery

import (
	godiscovery "github.com/fananchong/go-discovery"
	proto "github.com/golang/protobuf/proto"
)

type Node struct {
	godiscovery.Node
	Info    ServerInfo
	Servers IServers
}

func (this *Node) InitPolicy(policy Policy) {
	switch policy {
	case Ordered:
		this.Servers = NewServersPolicyOrdered()
	case Random:
		this.Servers = NewServersPolicyRandom()
	case RoundRobin:
		this.Servers = NewServersPolicyRoundRobin()
	default:
	}
}

func (this *Node) GetBase() interface{} {
	return this
}

func (this *Node) OnNodeUpdate(nodeType int, id string, data []byte) {
	info := &ServerInfo{}
	err := proto.Unmarshal(data, info)
	if err == nil {
		this.Servers.Set(nodeType, id, info)
	} else {
		this.GetLogger().Errorln("[NODE] DATA ERROR!")
	}
}

func (this *Node) OnNodeJoin(nodeType int, id string, data []byte) {
	info := &ServerInfo{}
	err := proto.Unmarshal(data, info)
	if err == nil {
		this.Servers.Set(nodeType, id, info)
	} else {
		this.GetLogger().Errorln("[NODE] DATA ERROR!")
	}
}

func (this *Node) OnNodeLeave(nodeType int, id string) {
	this.Servers.Delete(nodeType, id)
}

func (this *Node) GetPutData() (string, error) {
	info := this.Info
	data, err := proto.Marshal(&info)
	return string(data), err
}
