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

func GetNamespace(t int) string {
	if v, ok := gNodeTypes[t]; ok {
		return v.ns
	}
	return ""
}

func GetServiceName(t int) string {
	if v, ok := gNodeTypes[t]; ok {
		return v.svc
	}
	return ""
}

// 给 k8s_api_windows.go 使用，不需要考虑效率问题
func GetServiceType(ns, svc string) int {
	for k, v := range gNodeTypes {
		if v.ns == ns && v.svc == svc {
			return k
		}
	}
	return 0
}
