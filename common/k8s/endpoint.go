package k8s

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
