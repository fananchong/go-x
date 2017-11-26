package gonet

import (
	"net"
	"time"
)

type TcpClient struct {
	TcpTask
}

func (this *TcpClient) Connect(address string, tcpTask ITcpTask) bool {
	if this.TcpTask.Conn != nil {
		xlog.Warningln("[连接] 该地址上的连接已经存在，将会被销毁 ", address)
		this.TcpTask.Close()
	}

	conn, err := this.connectDetail(address)
	if err == nil {
		this.TcpTask = *NewTcpTask(conn, tcpTask)
		this.TcpTask.Start()
		return true
	} else {
		return false
	}
}

func (this *TcpClient) connectDetail(address string) (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		xlog.Errorln("[连接] ", err)
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		xlog.Errorln("[连接] ", err)
		return nil, err
	}

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(1 * time.Minute)
	conn.SetNoDelay(true)
	conn.SetWriteBuffer(128 * 1024)
	conn.SetReadBuffer(128 * 1024)

	return conn, nil
}
