package discovery

import "github.com/fananchong/gomap"

type ServersPolicyRandom struct {
	ss map[int]*gomap.RandomMap
}

func NewServersPolicyRandom() *ServersPolicyRandom {
	return &ServersPolicyRandom{ss: make(map[int]*gomap.RandomMap)}
}

func (this *ServersPolicyRandom) GetOne(nodeType int) (*ServerInfo, bool) {
	if ss, ok := this.ss[nodeType]; ok {
		if _, info, ok2 := ss.Random(); ok2 {
			return info.(*ServerInfo), true
		}
	}
	return nil, false
}
func (this *ServersPolicyRandom) GetAll(nodeType int) ([]*ServerInfo, bool) {
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
func (this *ServersPolicyRandom) Set(nodeType int, id string, val *ServerInfo) {
	if ss, ok := this.ss[nodeType]; ok {
		ss.Set(id, val)
	} else {
		ss := gomap.NewRandomMap()
		ss.Set(id, val)
		this.ss[nodeType] = ss
	}
}

func (this *ServersPolicyRandom) Delete(nodeType int, id string) {
	if ss, ok := this.ss[nodeType]; ok {
		ss.Delete(id)
	}
}
