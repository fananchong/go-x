package gotcp

import (
	"fmt"
	"net"
)

func GetVaildPort() int {
	port := 10000
	for {
		port = port + 1
		address := fmt.Sprintf(":%d", port)
		tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
		if err != nil {
			xlog.Errorln(err)
			continue
		}
		listener, err := net.ListenTCP("tcp", tcpAddr)
		if err != nil {
			xlog.Errorln(err)
			if listener != nil {
				listener.Close()
			}
			continue
		}
		listener.Close()
		return port
	}
	return 0
}
