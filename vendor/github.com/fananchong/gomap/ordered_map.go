package gomap

import (
	"sort"
)

type OrderedMap struct {
	MapBase
	less func(v1, v2 interface{}) bool
}

func NewOrderedMap(less func(v1, v2 interface{}) bool) *OrderedMap {
	this := &OrderedMap{less: less}
	this.Init()
	return this
}

func (this *OrderedMap) Set(key interface{}, val interface{}) bool {
	needSort := this.MapBase.Set(key, val)
	if needSort {
		sort.Sort(this)
	}
	return needSort
}

func (this *OrderedMap) First() (key, val interface{}, ok bool) {
	if len(this.s) <= 0 {
		return nil, nil, false
	}
	elem := this.s[0]
	return elem.key, elem.val, true
}

func (this *OrderedMap) Back() (key, val interface{}, ok bool) {
	len := len(this.s)
	if len <= 0 {
		return nil, nil, false
	}
	elem := this.s[len-1]
	return elem.key, elem.val, true
}

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
