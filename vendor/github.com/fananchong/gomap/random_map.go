package gomap

import "math/rand"

type RandomMap struct {
	MapBase
}

func NewRandomMap() *RandomMap {
	this := &RandomMap{}
	this.Init()
	return this
}

func (this *RandomMap) Random() (key, val interface{}, ok bool) {
	len := len(this.s)
	if len <= 0 {
		return nil, nil, false
	}
	index := rand.Intn(len)
	return this.s[index].key, this.s[index].val, true
}
