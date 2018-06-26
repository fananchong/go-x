package discovery

import (
	"sync"

	"github.com/fananchong/gomap"
)

type IServers interface {
	GetOne(nodeType int) (uint32, *ServerInfo, bool)
	GetAll(nodeType int) ([]*ServerInfo, bool)
	Set(nodeType int, id uint32, val *ServerInfo)
	Delete(nodeType int, id uint32)
	GetByID(id uint32) (*ServerInfo, bool)
	Count(nodeType int) int
}

type IMap interface {
	GetOne() (key, val interface{}, ok bool)

	Set(key interface{}, val interface{}) bool
	Delete(key interface{})
	Iterator() *gomap.Iterator
	Count() int
}

type Servers struct {
	ss      map[int]IMap
	creator func() IMap
	mutex   sync.RWMutex
	ssByID  map[uint32]*ServerInfo
}

func NewServers(creator func() IMap) *Servers {
	return &Servers{
		ss:      make(map[int]IMap),
		creator: creator,
		ssByID:  make(map[uint32]*ServerInfo),
	}
}

func (this *Servers) GetOne(nodeType int) (uint32, *ServerInfo, bool) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	if m, ok := this.ss[nodeType]; ok {
		if key, info, ok := m.GetOne(); ok {
			return key.(uint32), info.(*ServerInfo), true
		}
	}
	return 0, nil, false
}

func (this *Servers) GetAll(nodeType int) ([]*ServerInfo, bool) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
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
func (this *Servers) Set(nodeType int, id uint32, val *ServerInfo) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if m, ok := this.ss[nodeType]; ok {
		m.Set(id, val)
	} else {
		m := this.creator()
		m.Set(id, val)
		this.ss[nodeType] = m
	}
	this.ssByID[id] = val
}

func (this *Servers) Delete(nodeType int, id uint32) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if m, ok := this.ss[nodeType]; ok {
		m.Delete(id)
	}
	if _, ok := this.ssByID[id]; ok {
		delete(this.ssByID, id)
	}
}
func (this *Servers) GetByID(id uint32) (m *ServerInfo, ok bool) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	m, ok = this.ssByID[id]
	return
}

func (this *Servers) Count(nodeType int) int {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	if m, ok := this.ss[nodeType]; ok {
		return m.Count()
	}
	return 0
}
