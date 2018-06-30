package k8s

const ID_BASE = 100000

type Endpoint struct {
	NodeType int
	Index    int
	IP       string
	Ports    map[string]int
}

func NewEndpoint() *Endpoint {
	return &Endpoint{
		Ports: make(map[string]int),
	}
}

func (this *Endpoint) Id() uint32 {
	return uint32(this.NodeType*ID_BASE + this.Index)
}
