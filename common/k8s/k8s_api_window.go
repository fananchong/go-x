package k8s

func GetEndpoints(namespace, service string) ([]*Endpoint, error) {
	item := NewEndpoint()
	item.Index = 0
	item.IP = "localhost"

	switch service {
	case "gateway":
		item.NodeType = 2
		item.Ports[""] = 30200
	case "hub":
		item.NodeType = 5
		item.Ports[""] = 30500
	}

	var ips []*Endpoint
	ips = append(ips, item)
	return ips, nil
}

func getIndex(service string, name string) int {
	return 0
}
