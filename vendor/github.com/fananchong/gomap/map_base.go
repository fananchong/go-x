package gomap

type MapBase struct {
	m map[interface{}]*Element
	s []*Element
}

func (this *MapBase) Init() {
	this.m = make(map[interface{}]*Element)
	this.s = make([]*Element, 0)
}

func (this *MapBase) Get(key interface{}) (interface{}, bool) {
	if elem, ok := this.m[key]; ok {
		if elem.key != key {
			panic("[MapBase.Get] DATA ERROR!")
		}
		return elem.val, true
	}
	return nil, false
}

func (this *MapBase) Set(key interface{}, val interface{}) bool {
	change := true
	if elem, ok := this.m[key]; ok {
		if elem.key != key {
			panic("[MapBase.Set] DATA ERROR!")
		}
		if elem.val != val {
			elem.val = val
		} else {
			change = false
		}
	} else {
		elem := &Element{key, val, len(this.s)}
		this.m[key] = elem
		this.s = append(this.s, elem)
	}
	return change
}

func (this *MapBase) Delete(key interface{}) {
	if elem, ok := this.m[key]; ok {
		if elem.key != key {
			panic("[MapBase.Delete] DATA ERROR!")
		}
		if elem.pos < len(this.s) {
			this.s = append(this.s[:elem.pos], this.s[elem.pos+1:]...)
			for i := elem.pos + 1; i < len(this.s); i++ {
				this.s[i].pos = i
			}
		}
		delete(this.m, key)
	}
}

func (this *MapBase) Iterator() *Iterator {
	return &Iterator{
		data:  this.s,
		index: 0,
	}
}

func (this *MapBase) Count() int {
	return len(this.s)
}
