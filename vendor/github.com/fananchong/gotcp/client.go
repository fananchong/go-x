package gotcp

import (
	"net"
	"time"
)

func (this *Session) Connect(address string, derived ISession) bool {
	if this.IsClosed() == false {
		xlog.Errorln("close session. server address =", this.RemoteAddr())
		this.Close()
	}
	conn, err := connectDetail(address)
	if err == nil {
		this.Init(conn, nil, derived)
		this.Start()
		xlog.Infoln("connect server success. server address =", this.RemoteAddr())
		return true
	} else {
		xlog.Errorln(err)
		return false
	}
}

func connectDetail(address string) (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(1 * time.Minute)
	conn.SetNoDelay(true)
	conn.SetWriteBuffer(128 * 1024)
	conn.SetReadBuffer(128 * 1024)
	return conn, nil
}
