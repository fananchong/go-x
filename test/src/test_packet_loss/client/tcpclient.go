package main

import (
	"flag"
	"github.com/fananchong/gonet"
	"github.com/golang/glog"
	"io"
	"net"
	"time"
)

func TcpClient() {
	faddr := flag.Lookup("tcpaddr")
	addr, _ := net.ResolveTCPAddr("tcp4", faddr.Value.String())
	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		panic(err)
	}

	glog.Infoln("connect to ", conn.RemoteAddr().String())

	go func() {

		conn.SetNoDelay(true)
		conn.SetWriteBuffer(128 * 1024)
		conn.SetReadBuffer(128 * 1024)

		buf := gonet.NewByteBuffer()
		msglen := len(MSG)
		var tempbuf [1024]byte
		for {
			leastlen := msglen - buf.RdSize()
			readnum, err := io.ReadAtLeast(conn, tempbuf[0:], leastlen)
			if err != nil {
				glog.Infoln(err)
				return
			}
			buf.Append(tempbuf[:readnum])
			now := time.Now().UnixNano()
			for buf.RdSize() >= msglen {
				msgbuff := buf.RdBuf()
				onTcpRecv(msgbuff[:msglen], now)
				buf.RdFlip(msglen)
			}
		}
	}()
}

var (
	preTCPRecvTime int64 = 0
)

func onTcpRecv(data []byte, now int64) {
	if len(data) != len(MSG) {
		panic("data len error!")
	}
	for i := 0; i < len(data); i++ {
		if data[i] != MSG[i] {
			panic("data error!")
		}
	}

	if preTCPRecvTime == 0 {
		preTCPRecvTime = now
	}

	detal := (now - preTCPRecvTime) / int64(time.Millisecond)
	preTCPRecvTime = now

	g_chart.AddTcpData(detal)
}
