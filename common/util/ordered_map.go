package util

import (
	"sort"
)

type itemInfo struct {
	key interface{}
	val interface{}
	pos int
}

type OrderedMap struct {
	m    map[interface{}]*itemInfo
	s    []*itemInfo
	less func(v1, v2 interface{}) bool
}

func NewOrderedMap(less func(v1, v2 interface{}) bool) *OrderedMap {
	return &OrderedMap{
		m:    make(map[interface{}]*itemInfo),
		s:    make([]*itemInfo, 0),
		less: less,
	}
}

// First / Back / Get / Set / Delete

func (this *OrderedMap) First() (key, val interface{}, ok bool) {
	if len(this.s) <= 0 {
		return nil, nil, false
	}
	item := this.s[0]
	return item.key, item.val, true
}

func (this *OrderedMap) Back() (key, val interface{}, ok bool) {
	len := len(this.s)
	if len <= 0 {
		return nil, nil, false
	}
	item := this.s[len-1]
	return item.key, item.val, true
}

func (this *OrderedMap) Get(key interface{}) (interface{}, bool) {
	if item, ok := this.m[key]; ok {
		if item.key != key {
			panic("[OrderedMap.Get] DATA ERROR!")
		}
		return item.val, true
	}
	return nil, false
}

func (this *OrderedMap) Set(key interface{}, val interface{}) {
	needSort := true
	if item, ok := this.m[key]; ok {
		if item.key != key {
			panic("[OrderedMap.Set] DATA ERROR!")
		}
		if item.val != val {
			item.val = val
		} else {
			needSort = false
		}
	} else {
		item := &itemInfo{key, val, len(this.s)}
		this.m[key] = item
		this.s = append(this.s, item)
	}
	if needSort {
		sort.Sort(this)
	}
}

func (this *OrderedMap) Delete(key interface{}) {
	if item, ok := this.m[key]; ok {
		if item.key != key {
			panic("[OrderedMap.Delete] DATA ERROR!")
		}
		if item.pos < len(this.s) {
			this.s = append(this.s[:item.pos], this.s[item.pos+1:]...)
			for i := item.pos + 1; i < len(this.s); i++ {
				this.s[i].pos = i
			}
		}
		delete(this.m, item)
	}
}

// Iterate

func (this *OrderedMap) Iterator() *Iterator {
	return &Iterator{
		data:  this.s,
		index: 0,
	}
}

type Iterator struct {
	data  []*itemInfo
	index int
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.data)
}

func (i *Iterator) Next() (key, val interface{}) {
	tempindex := i.index
	i.index++
	return i.data[tempindex].key, i.data[tempindex].val
}

// sort interface

func (this *OrderedMap) Len() int {
	return len(this.s)
}

func (this *OrderedMap) Swap(i, j int) {
	tempk := this.s[i].key
	tempv := this.s[i].val
	this.s[i].pos = i
	this.s[i].key = this.s[j].key
	this.s[i].val = this.s[j].val
	this.s[j].pos = j
	this.s[j].key = tempk
	this.s[j].val = tempv
	this.m[this.s[i].key] = this.s[i]
	this.m[this.s[j].key] = this.s[j]
}

func (this *OrderedMap) Less(i, j int) bool {
	return this.less(this.s[i].val, this.s[j].val)
}
