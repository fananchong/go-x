package k8s

type NodeType struct {
	t   int
	ns  string
	svc string
}

var gNodeTypes map[int]*NodeType = make(map[int]*NodeType)

func RegisterNodeType(t int, ns, svc string) {
	nt := &NodeType{
		t:   t,
		ns:  ns,
		svc: svc,
	}
	gNodeTypes[t] = nt
}
