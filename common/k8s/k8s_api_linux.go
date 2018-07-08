package k8s

import (
	"context"
	"strconv"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

func GetEndpoints(namespace, service string) ([]*Endpoint, error) {
	client, err := k8s.NewInClusterClient()
	if err != nil {
		return nil, err
	}

	var ips []*Endpoint
	var endpoints corev1.Endpoints
	err = client.Get(context.Background(), namespace, service, &endpoints)
	if err != nil {
		return nil, err
	}

	for _, endpoint := range endpoints.Subsets {
		for _, address := range endpoint.Addresses {
			index := getIndex(service, *address.Hostname)
			item := NewEndpoint()
			item.IP = *address.Ip
			item.Index = index
			for _, port := range endpoint.Ports {
				item.Ports[*port.Name] = int(*port.Port) + index
			}
			ips = append(ips, item)
		}
	}

	return ips, nil
}

func getIndex(service string, name string) int {
	if len(service) >= len(name) {
		return 0
	}
	id, err := strconv.Atoi(name[len(service)+1:])
	if err != nil {
		return 0
	}
	return id
}
