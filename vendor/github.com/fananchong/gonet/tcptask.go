package gonet

import (
	"io"
	"net"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

type ITcpTask interface {
	ParseMsg(data []byte, flag byte)
	OnClose()
}

const (
	cmd_max_size    = 128 * 1024
	cmd_header_size = 4 // 3字节指令长度 1字节是否压缩
	cmd_verify_time = 10
	send_chan_size  = 1024
)

type TcpTask struct {
	closed     int32
	verified   bool
	stopedChan chan struct{}
	recvBuff   *ByteBuffer
	sendChan   chan []byte
	Conn       net.Conn
	Derived    ITcpTask
}

func NewTcpTask(conn net.Conn, tcpTask ITcpTask) *TcpTask {
	task := &TcpTask{
		closed:     -1,
		verified:   false,
		Conn:       conn,
		Derived:    tcpTask,
		stopedChan: make(chan struct{}),
		recvBuff:   NewByteBuffer(),
		sendChan:   make(chan []byte, send_chan_size),
	}
	return task
}

func (this *TcpTask) RemoteAddr() string {
	if this.Conn == nil {
		return ""
	}
	return this.Conn.RemoteAddr().String()
}

func (this *TcpTask) Start() {
	if atomic.CompareAndSwapInt32(&this.closed, -1, 0) {
		job := &sync.WaitGroup{}
		job.Add(2)
		go this.sendloop(job)
		go this.recvloop(job)
		job.Wait()
		xlog.Infoln("[连接] 收到连接 ", this.RemoteAddr())
	}
}

func (this *TcpTask) Close() {
	if atomic.CompareAndSwapInt32(&this.closed, 0, 1) {
		xlog.Infoln("[连接] 断开连接 ", this.RemoteAddr())
		this.Conn.Close()
		close(this.stopedChan)
		close(this.sendChan)
		this.recvBuff.Reset()
		this.Derived.OnClose()
		this.Derived = nil
	}
}

func (this *TcpTask) IsClosed() bool {
	return atomic.LoadInt32(&this.closed) != 0
}

func (this *TcpTask) Verify() {
	this.verified = true
}

func (this *TcpTask) IsVerified() bool {
	return this.verified
}

func (this *TcpTask) AsyncSend(buffer []byte, flag byte) bool {
	if this.IsClosed() {
		return false
	}
	bsize := len(buffer)
	data := []byte{byte(bsize), byte(bsize >> 8), byte(bsize >> 16), flag}
	if bsize != 0 {
		data = append(data, buffer...)
	}
	select {
	case this.sendChan <- data:
	default:
		xlog.Errorln("发送缓冲区满！连接将被关闭！")
		this.Close()
		return false
	}
	return true
}

func (this *TcpTask) recvloop(job *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			xlog.Errorln("[异常] ", err, "\n", string(debug.Stack()))
		}
	}()
	defer this.Close()

	var (
		neednum   int
		readnum   int
		err       error
		totalsize int
		datasize  int
		msgbuff   []byte
	)

	job.Done()

	for {
		totalsize = this.recvBuff.RdSize()

		if totalsize < cmd_header_size {

			neednum = cmd_header_size - totalsize
			this.recvBuff.WrGrow(neednum)

			readnum, err = io.ReadAtLeast(this.Conn, this.recvBuff.WrBuf(), neednum)
			if err != nil {
				xlog.Infoln("[连接] 接收失败 ", this.RemoteAddr(), ",", err)
				return
			}

			this.recvBuff.WrFlip(readnum)
			totalsize = this.recvBuff.RdSize()
		}

		msgbuff = this.recvBuff.RdBuf()

		datasize = int(msgbuff[0]) | int(msgbuff[1])<<8 | int(msgbuff[2])<<16
		if datasize > cmd_max_size-cmd_header_size {
			xlog.Errorln("[连接] 数据超过最大值 ", this.RemoteAddr(), ",", datasize)
			return
		}
		if datasize <= 0 {
			xlog.Errorln("[连接] 数据长度为0或负值", this.RemoteAddr(), ",", datasize)
			return
		}

		if totalsize < cmd_header_size+datasize {

			neednum = cmd_header_size + datasize - totalsize
			this.recvBuff.WrGrow(neednum)

			readnum, err = io.ReadAtLeast(this.Conn, this.recvBuff.WrBuf(), neednum)
			if err != nil {
				xlog.Infoln("[连接] 接收失败 ", this.RemoteAddr(), ",", err)
				return
			}

			this.recvBuff.WrFlip(readnum)
			msgbuff = this.recvBuff.RdBuf()
		}

		this.Derived.ParseMsg(msgbuff[cmd_header_size:cmd_header_size+datasize], msgbuff[3])
		this.recvBuff.RdFlip(cmd_header_size + datasize)
	}
}

func (this *TcpTask) sendloop(job *sync.WaitGroup) {
	var (
		tmpByte  = NewByteBuffer()
		timeout  = time.NewTimer(time.Second * cmd_verify_time)
		writenum int
		err      error
	)

	defer func() {
		if err := recover(); err != nil {
			xlog.Errorln("[异常] ", err, "\n", string(debug.Stack()))
		}
		this.Close()
		timeout.Stop()
	}()

	job.Done()

	for {
		select {
		case buff := <-this.sendChan:
			tmpByte.Append(buff)
			for {
				if !tmpByte.RdReady() {
					tmpByte.Reset()
					break
				}
				writenum, err = this.Conn.Write(tmpByte.RdBuf()[:tmpByte.RdSize()])
				if err != nil {
					xlog.Infoln("[连接] 发送失败 ", this.RemoteAddr(), ",", err)
					return
				}
				tmpByte.RdFlip(writenum)
			}
		case <-this.stopedChan:
			return
		case <-timeout.C:
			if !this.IsVerified() {
				xlog.Infoln("[连接] 验证超时 ", this.RemoteAddr())
				return
			}
		}
	}
}
