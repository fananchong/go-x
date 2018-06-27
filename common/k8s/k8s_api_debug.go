// +build debug

package k8s

func GetEndpoints(namespace, service string) ([]*Endpoint, error) {

	// 仅举例，可以从配置文件中读取

	item := NewEndpoint()
	item.Index = 0
	item.IP = "localhost"
	item.Ports[""] = 3000
	var ips []*Endpoint
	ips = append(ips, item)
	return ips, nil
}

func GetVaildPort(namespace, service string) (map[string]int, error) {

	// 仅举例，可以从配置文件中读取

	ports := make(map[string]int)
	ports[""] = 3000
	return ports, nil
}
