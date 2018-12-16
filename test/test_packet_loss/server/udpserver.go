package main

import (
	"net"
	"time"
)

func UdpServer() {
	addr, err := net.ResolveUDPAddr("udp", ":3001")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	conn.SetWriteBuffer(128 * 1024)
	conn.SetReadBuffer(128 * 1024)

	var buf [20]byte
	_, raddr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		base.XLOG.Infoln(err)
		return
	}

	// 每100ms发送一次 hello消息
	t := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-t.C:
			conn.WriteToUDP(MSG, raddr)
		}
	}
}
