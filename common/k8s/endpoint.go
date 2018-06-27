package k8s

type Endpoint struct {
	Index int
	IP    string
	Ports map[string]int
}

func NewEndpoint() *Endpoint {
	return &Endpoint{
		Ports: make(map[string]int),
	}
}
