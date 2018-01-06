package discovery

import (
	"sync"

	"github.com/fananchong/gomap"
)

type IServers interface {
	GetOne(nodeType int) (*ServerInfo, bool)
	GetAll(nodeType int) ([]*ServerInfo, bool)
	Set(nodeType int, id string, val *ServerInfo)
	Delete(nodeType int, id string)
}

type IMap interface {
	GetOne() (key, val interface{}, ok bool)

	Set(key interface{}, val interface{}) bool
	Delete(key interface{})
	Iterator() *gomap.Iterator
}

type Servers struct {
	ss      map[int]IMap
	creator func() IMap
	mutex   sync.Mutex
}

func (this *Servers) GetOne(nodeType int) (*ServerInfo, bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if m, ok := this.ss[nodeType]; ok {
		if _, info, ok := m.GetOne(); ok {
			return info.(*ServerInfo), true
		}
	}
	return nil, false
}

func (this *Servers) GetAll(nodeType int) ([]*ServerInfo, bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if m, ok := this.ss[nodeType]; ok {
		ret := make([]*ServerInfo, 0)
		for iter := m.Iterator(); iter.HasNext(); {
			_, s := iter.Next()
			ret = append(ret, s.(*ServerInfo))
		}
		return ret, true
	}
	return nil, false
}
func (this *Servers) Set(nodeType int, id string, val *ServerInfo) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if m, ok := this.ss[nodeType]; ok {
		m.Set(id, val)
	} else {
		m := this.creator()
		m.Set(id, val)
		this.ss[nodeType] = m
	}
}

func (this *Servers) Delete(nodeType int, id string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if m, ok := this.ss[nodeType]; ok {
		m.Delete(id)
	}
}
