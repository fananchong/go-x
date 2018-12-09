package main

import (
	"net"
	"time"
)

func TcpServer() {
	addr, _ := net.ResolveTCPAddr("tcp", ":3000")
	lis, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.AcceptTCP()
		if err != nil {
			xlog.Infoln(err)
			continue
		}

		xlog.Infoln("on connect. addr =", conn.RemoteAddr())

		go func() {
			conn.SetNoDelay(true)
			conn.SetWriteBuffer(128 * 1024)
			conn.SetReadBuffer(128 * 1024)

			// 每100ms发送一次 hello消息
			t := time.NewTicker(100 * time.Millisecond)
			for {
				select {
				case <-t.C:
					conn.Write(MSG)
				}
			}
		}()
	}
}
