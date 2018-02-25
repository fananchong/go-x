package gomap

type Iterator struct {
	data  []*Element
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
