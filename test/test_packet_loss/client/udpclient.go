package main

import (
	"flag"
	"net"
	"time"
)

func UdpClient() {
	faddr := flag.Lookup("udpaddr")
	addr, _ := net.ResolveUDPAddr("udp", faddr.Value.String())
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}

	base.XLOG.Infoln("connect to ", conn.RemoteAddr().String())

	go func() {

		conn.SetWriteBuffer(128 * 1024)
		conn.SetReadBuffer(128 * 1024)
		conn.Write(MSG)

		for {
			var tempbuf [1024]byte
			readnum, err := conn.Read(tempbuf[0:])
			if err != nil {
				base.XLOG.Infoln(err)
				return
			}
			now := time.Now().UnixNano()
			onUdpRecv(tempbuf[:readnum], now)
		}
	}()
}

var (
	preUDPRecvTime int64 = 0
)

func onUdpRecv(data []byte, now int64) {
	if len(data) != len(MSG) {
		panic("data len error!")
	}
	for i := 0; i < len(data); i++ {
		if data[i] != MSG[i] {
			panic("data error!")
		}
	}

	if preUDPRecvTime == 0 {
		preUDPRecvTime = now
	}

	detal := (now - preUDPRecvTime) / int64(time.Millisecond)
	preUDPRecvTime = now

	g_chart.AddUdpData(detal)
}
