package discovery

import "github.com/fananchong/gomap"

type ServersPolicyRoundRobin struct {
	ss map[int]*gomap.RoundRobinMap
}

func NewServersPolicyRoundRobin() *ServersPolicyRoundRobin {
	return &ServersPolicyRoundRobin{ss: make(map[int]*gomap.RoundRobinMap)}
}

func (this *ServersPolicyRoundRobin) GetOne(nodeType int) (*ServerInfo, bool) {
	if ss, ok := this.ss[nodeType]; ok {
		if _, info, ok2 := ss.RoundRobin(); ok2 {
			return info.(*ServerInfo), true
		}
	}
	return nil, false
}
func (this *ServersPolicyRoundRobin) GetAll(nodeType int) ([]*ServerInfo, bool) {
	if ss, ok := this.ss[nodeType]; ok {
		ret := make([]*ServerInfo, 0)
		for iter := ss.Iterator(); iter.HasNext(); {
			_, s := iter.Next()
			ret = append(ret, s.(*ServerInfo))
		}
		return ret, true
	}
	return nil, false
}
func (this *ServersPolicyRoundRobin) Set(nodeType int, id string, val *ServerInfo) {
	if ss, ok := this.ss[nodeType]; ok {
		ss.Set(id, val)
	} else {
		ss := gomap.NewRoundRobinMap()
		ss.Set(id, val)
		this.ss[nodeType] = ss
	}
}

func (this *ServersPolicyRoundRobin) Delete(nodeType int, id string) {
	if ss, ok := this.ss[nodeType]; ok {
		ss.Delete(id)
	}
}
