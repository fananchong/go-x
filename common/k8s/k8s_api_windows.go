package k8s

func GetEndpoints(namespace, service string) ([]*Endpoint, error) {
	item := NewEndpoint()
	item.Index = 0
	item.IP = "localhost"

	var ips []*Endpoint
	t := GetServiceType(namespace, service)
	if t > 0 {

		item.NodeType = t
		item.Ports[""] = 30000 + t*100
		ips = append(ips, item)
		return ips, nil
	}
	return ips, nil
}

func getIndex(service string, name string) int {
	return 0
}
