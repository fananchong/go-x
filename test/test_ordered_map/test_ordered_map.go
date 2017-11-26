package main

import (
	"fmt"

	"github.com/fananchong/go-x/common/util"
)

func less(v1, v2 interface{}) bool {
	return v1.(int) < v2.(int)
}

func main() {
	// Init new OrderedMap
	om := util.NewOrderedMap(less)

	// Set key
	om.Set("a", 1)
	om.Set("c", 3)
	om.Set("b", 2)
	om.Set("d", 4)
	om.Set("e", 2)
	om.Set("f", 2)

	// Same interface as builtin map
	if val, ok := om.Get("a"); ok == true {
		// Found key "a"
		fmt.Println(val)
	}

	// Delete a key
	om.Delete("e")

	// Failed Get lookup becase we deleted "c"
	if _, ok := om.Get("e"); ok == false {
		// Did not find key "e"
		fmt.Println("e not found")
	}

	// Iterator
	for it := om.Iterator(); it.HasNext(); {
		fmt.Println(it.Next())
	}
}
