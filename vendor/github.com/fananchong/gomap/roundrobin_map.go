package gomap

type RoundRobinMap struct {
	MapBase
	curIndex int
}

func NewRoundRobinMap() *RoundRobinMap {
	this := &RoundRobinMap{}
	this.Init()
	return this
}

func (this *RoundRobinMap) RoundRobin() (key, val interface{}, ok bool) {
	len := len(this.s)
	if len <= 0 {
		return nil, nil, false
	}
	if this.curIndex >= len {
		this.curIndex = 0
	}
	index := this.curIndex
	this.curIndex++
	return this.s[index].key, this.s[index].val, true
}
