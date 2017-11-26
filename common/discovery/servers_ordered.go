package discovery

import "github.com/fananchong/gomap"

type ServersPolicyOrdered struct {
	ss map[int]*gomap.OrderedMap
}

func NewServersPolicyOrdered() *ServersPolicyOrdered {
	return &ServersPolicyOrdered{ss: make(map[int]*gomap.OrderedMap)}
}

func (this *ServersPolicyOrdered) GetOne(nodeType int) (*ServerInfo, bool) {
	if ss, ok := this.ss[nodeType]; ok {
		if _, info, ok2 := ss.Back(); ok2 {
			return info.(*ServerInfo), true
		}
	}
	return nil, false
}
func (this *ServersPolicyOrdered) GetAll(nodeType int) ([]*ServerInfo, bool) {
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
func (this *ServersPolicyOrdered) Set(nodeType int, id string, val *ServerInfo) {
	if ss, ok := this.ss[nodeType]; ok {
		ss.Set(id, val)
	} else {
		ss := gomap.NewOrderedMap(less)
		ss.Set(id, val)
		this.ss[nodeType] = ss
	}
}

func (this *ServersPolicyOrdered) Delete(nodeType int, id string) {
	if ss, ok := this.ss[nodeType]; ok {
		ss.Delete(id)
	}
}

func less(v1, v2 interface{}) bool {
	return v1.(*ServerInfo).Ordered >= v2.(*ServerInfo).Ordered
}
