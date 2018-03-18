package gotcp

import (
	"context"
	"net"
	"reflect"
	"time"
)

type Server struct {
	listener  *net.TCPListener
	sessType  reflect.Type
	ctx       context.Context
	ctxCancel context.CancelFunc
	address   string
}

func (this *Server) RegisterSessType(v interface{}) {
	this.sessType = reflect.ValueOf(v).Type()
}

func (this *Server) Start(address string) bool {
	this.address = address
	if this.listener != nil {
		return true
	}
	err := this.bind(address)
	if err != nil {
		xlog.Errorln(err)
		return false
	}
	this.ctx, this.ctxCancel = context.WithCancel(context.Background())
	go this.loop()
	return true
}

func (this *Server) Close() {
	this.ctxCancel()
	this.listener.Close()
	this.listener = nil
}

func (this *Server) loop() {
	for {
		select {
		case <-this.ctx.Done():
			xlog.Infoln("server close. address =", this.address)
			return
		default:
			conn, err := this.accept()
			if err == nil {
				sess := reflect.New(this.sessType)
				f := sess.MethodByName("Init")
				f.Call([]reflect.Value{reflect.ValueOf(conn), reflect.ValueOf(this.ctx), sess})
				f = sess.MethodByName("Start")
				f.Call([]reflect.Value{})
				f = sess.MethodByName("RemoteAddr")
				addr := f.Call([]reflect.Value{})
				xlog.Infoln("connect come in. client address =", addr)
			}
		}
	}
}

func (this *Server) bind(address string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}
	this.listener = listener
	return nil
}

func (this *Server) accept() (*net.TCPConn, error) {
	conn, err := this.listener.AcceptTCP()
	if err != nil {
		if opErr, ok := err.(*net.OpError); ok && !opErr.Timeout() {
			xlog.Errorln(err)
		}
		return nil, err
	}
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(1 * time.Minute)
	conn.SetNoDelay(true)
	conn.SetWriteBuffer(128 * 1024)
	conn.SetReadBuffer(128 * 1024)
	return conn, nil
}
