package service

import (
	"context"
	"net"
	"sync"

	"github.com/fananchong/go-x/base"
	"github.com/fananchong/go-x/common_services/proto"
	discovery "github.com/fananchong/go-x/internal/common/k8s/serverlist"
	"github.com/fananchong/gotcp"
	proto1 "github.com/golang/protobuf/proto"
)

type SessionIntranet struct {
	gotcp.Session
	Id             uint32
	Msgs           *Messages
	DefaultHandler func(uid uint64, data []byte, flag byte)
}

func (this *SessionIntranet) Init(conn net.Conn, root context.Context, derived gotcp.ISession) {
	this.Session.Init(conn, root, derived)
	this.Msgs = NewMessages()
}

func (this *SessionIntranet) OnRecv(data []byte, flag byte) {
	if this.IsVerified() == false {
		this.doVerify(data, flag)
		return
	}
	var uid uint64 = 0
LABEL_AGAIN:
	cmd := proto.MsgTypeCmd(gotcp.GetCmd(data))
	if handler, ok := this.Msgs.Handlers[uint32(cmd)]; ok {
		handler(uid, data, flag)
	} else {
		if uid == 0 && cmd == proto.MsgTypeCmd_Forward {
			msg := &proto.MsgForward{}
			if gotcp.DecodeCmd(data, flag, msg) == nil {
				base.XLOG.Errorln("decodeMsg fail.")
				return
			}
			data = msg.GetData()
			flag = byte(msg.GetFlag())
			uid = msg.GetUID()
			goto LABEL_AGAIN
		}

		if this.DefaultHandler != nil {
			this.DefaultHandler(uid, data, flag)
		}
	}
}

func (this *SessionIntranet) doVerify(data []byte, flag byte) {
	msg := &proto.MsgVerifyS{}
	if gotcp.DecodeCmd(data, flag, msg) == nil {
		base.XLOG.Errorln("decodeMsg fail.")
		this.Close()
		return
	}
	if msg.GetToken() != base.XARGS.Common.IntranetToken {
		base.XLOG.Errorln("token error.")
		this.Close()
	}
	this.Id = msg.GetId()
	xnodes.Store(this.Id, this)
	this.Verify()
	base.XLOG.Debugln("Id:", msg.GetId(), "verify success. My Node Id:", discovery.GetNode().Id())

	msg.Reset()
	msg.Id = discovery.GetNode().Id()
	msg.Token = base.XARGS.Common.IntranetToken
	this.SendMsg(uint64(proto.MsgTypeCmd_Verify), msg)
}

func (this *SessionIntranet) OnClose() {
	if _, loaded := xnodes.Load(this.Id); loaded {
		xnodes.Delete(this.Id)
	}
}

func (this *SessionIntranet) SendMsgtoClient(uid uint64, cmd uint64, msg proto1.Message) {
	data, flag, err := gotcp.EncodeCmd(cmd, msg)
	if err != nil {
		base.XLOG.Errorln(err)
		return
	}
	newMsg := &proto.MsgForward{}
	newMsg.UID = uid
	newMsg.Data = data
	newMsg.Flag = int32(flag)
	this.SendMsg(uint64(proto.MsgTypeCmd_Forward), newMsg)
}

func (this *SessionIntranet) Forward(id string, data []byte, flag byte) {
	if n, ok := xnodes.Load(id); ok {
		n.(*SessionIntranet).Send(data, flag)
	}
}

func (this *SessionIntranet) ForwardMsg(id string, cmd uint64, msg proto1.Message) {
	if n, ok := xnodes.Load(id); ok {
		n.(*SessionIntranet).SendMsg(cmd, msg)
	}
}

func (this *SessionIntranet) Broadcast(data []byte, flag byte) {
	xnodes.Range(func(key interface{}, val interface{}) bool {
		val.(*SessionIntranet).Send(data, flag)
		return true
	})
}

func (this *SessionIntranet) BroadcastExcludeMe(data []byte, flag byte) {
	xnodes.Range(func(key interface{}, val interface{}) bool {
		if key.(uint32) != this.Id {
			val.(*SessionIntranet).Send(data, flag)
		}
		return true
	})
}

func (this *SessionIntranet) BroadcastMsg(cmd uint64, msg proto1.Message) {
	xnodes.Range(func(key interface{}, val interface{}) bool {
		val.(*SessionIntranet).SendMsg(cmd, msg)
		return true
	})
}

func (this *SessionIntranet) BroadcastMsgExcludeMe(cmd uint64, msg proto1.Message) {
	xnodes.Range(func(key interface{}, val interface{}) bool {
		if key.(uint32) != this.Id {
			val.(*SessionIntranet).SendMsg(cmd, msg)
		}
		return true
	})
}

// 这里内网session管理，没有做成单例管理类。
// 请不要模仿这种不好的习惯:)
var xnodes sync.Map

// Handlers
type MessagesFunc func(uid uint64, data []byte, flag byte)

type Messages struct {
	Handlers map[uint32]MessagesFunc
}

func NewMessages() *Messages {
	return &Messages{
		Handlers: make(map[uint32]MessagesFunc),
	}
}

func (this *Messages) RegisterMessage(cmd uint32, f MessagesFunc) {
	this.Handlers[cmd] = f
}
