package gotcp

import (
	"context"
	"io"
	"net"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

type ISession interface {
	OnRecv(data []byte, flag byte)
	OnClose()
}

const (
	cmd_max_size    = 128 * 1024 // 消息最大长度
	cmd_header_size = 4          // 3字节指令长度 1字节是否压缩
	cmd_verify_time = 10         // 连接验证超时时间
	send_chan_size  = 1024       // 发送缓冲区大小
)

type Session struct {
	Conn      net.Conn
	ctx       context.Context
	ctxCancel context.CancelFunc
	sendChan  chan []byte
	sendCount int32
	closed    int32
	verified  bool
	Derived   ISession
}

func (this *Session) Init(conn net.Conn, root context.Context, derived ISession) {
	this.Derived = derived
	this.Conn = conn
	if root == nil {
		this.ctx, this.ctxCancel = context.WithCancel(context.Background())
	} else {
		this.ctx, this.ctxCancel = context.WithCancel(root)
	}
	this.sendChan = make(chan []byte, send_chan_size)
	atomic.StoreInt32(&this.sendCount, 0)
	atomic.StoreInt32(&this.closed, 0)
	this.verified = false
}

func (this *Session) Start() {
	if atomic.CompareAndSwapInt32(&this.closed, 0, 1) {
		job := &sync.WaitGroup{}
		job.Add(2)
		go this.sendloop(job)
		go this.recvloop(job)
		job.Wait()
	}
}

func (this *Session) Close() {
	if atomic.CompareAndSwapInt32(&this.closed, 1, 2) {
		xlog.Infoln("disconnect. remote address =", this.RemoteAddr())
		this.ctxCancel()
		this.Conn.Close()
		close(this.sendChan)
		this.Derived.OnClose()
		this.Derived = nil
	}
}

func (this *Session) IsClosed() bool {
	return atomic.LoadInt32(&this.closed) != 1
}

func (this *Session) Verify() {
	this.verified = true
}

func (this *Session) IsVerified() bool {
	return this.verified
}

func (this *Session) Send(buffer []byte, flag byte) bool {
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
		atomic.AddInt32(&this.sendCount, 1)
	default:
		xlog.Errorln("send buffer is full! the connection will be closed!")
		this.Close()
		return false
	}
	return true
}

func (this *Session) recvloop(job *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			xlog.Errorln("[except] ", err, "\n", string(debug.Stack()))
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
		recvBuff  *ByteBuffer = NewByteBuffer()
		timeout               = time.NewTimer(time.Second * cmd_verify_time)
	)
	defer timeout.Stop()

	job.Done()

	for {
		select {
		case <-timeout.C:
			if !this.IsVerified() {
				xlog.Infoln("verify timeout, remote address =", this.RemoteAddr())
				return
			}
		case <-this.ctx.Done():
			return
		default:
			totalsize = recvBuff.RdSize()
			if totalsize < cmd_header_size {
				neednum = cmd_header_size - totalsize
				recvBuff.WrGrow(neednum)
				readnum, err = io.ReadAtLeast(this.Conn, recvBuff.WrBuf(), neednum)
				if err != nil {
					xlog.Infoln("recv data fail. error =", err)
					return
				}
				recvBuff.WrFlip(readnum)
				totalsize = recvBuff.RdSize()
			}
			msgbuff = recvBuff.RdBuf()
			datasize = int(msgbuff[0]) | int(msgbuff[1])<<8 | int(msgbuff[2])<<16
			if datasize > cmd_max_size-cmd_header_size {
				xlog.Errorln("data exceed the maximum. datasize =", datasize)
				return
			}
			if datasize <= 0 {
				xlog.Errorln("data length is 0 or negative. datasize =", datasize)
				return
			}
			if totalsize < cmd_header_size+datasize {
				neednum = cmd_header_size + datasize - totalsize
				recvBuff.WrGrow(neednum)
				readnum, err = io.ReadAtLeast(this.Conn, recvBuff.WrBuf(), neednum)
				if err != nil {
					xlog.Infoln("recv data fail. error =", err)
					return
				}
				recvBuff.WrFlip(readnum)
				msgbuff = recvBuff.RdBuf()
			}
			this.Derived.OnRecv(msgbuff[cmd_header_size:cmd_header_size+datasize], msgbuff[3])
			recvBuff.RdFlip(cmd_header_size + datasize)
		}
	}
}

func (this *Session) sendloop(job *sync.WaitGroup) {
	var (
		tmpByte  = NewByteBuffer()
		writenum int
		err      error
	)

	defer func() {
		if err := recover(); err != nil {
			xlog.Errorln("[except] ", err, "\n", string(debug.Stack()))
		}
		this.Close()
	}()

	job.Done()

	for {
		select {
		case buff := <-this.sendChan:
			tmpByte.Append(buff)
			if atomic.AddInt32(&this.sendCount, -1) <= 0 {
				for {
					if !tmpByte.RdReady() {
						tmpByte.Reset()
						break
					}
					writenum, err = this.Conn.Write(tmpByte.RdBuf()[:tmpByte.RdSize()])
					if err != nil {
						xlog.Infoln("send data fail. err =", err)
						return
					}
					tmpByte.RdFlip(writenum)
				}
			}
		case <-this.ctx.Done():
			return
		}
	}
}

func (this *Session) RemoteAddr() string {
	if this.Conn == nil {
		return ""
	}
	return this.Conn.RemoteAddr().String()
}
