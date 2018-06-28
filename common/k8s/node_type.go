package k8s

import (
	"sync"
)

type NodeType struct {
	t   int
	ns  string
	svc string
}

var gNodeTypes sync.Map

func RegisterNodeType(t int, ns, svc string) {
	nt := &NodeType{
		t:   t,
		ns:  ns,
		svc: svc,
	}
	gNodeTypes.Store(t, nt)
}
