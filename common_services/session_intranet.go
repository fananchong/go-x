package common

import (
	"sync"

	"github.com/fananchong/go-discovery/serverlist"
	"github.com/fananchong/go-x/common_services/proto"
	"github.com/fananchong/gotcp"
	proto1 "github.com/golang/protobuf/proto"
)

type SessionIntranet struct {
	gotcp.Session
	Id         uint32
	Cmds       map[uint64]func(data []byte, flag byte)
	DefaultCmd func(data []byte, flag byte)
}

func (this *SessionIntranet) OnRecv(data []byte, flag byte) {
	if this.IsVerified() == false {
		this.doVerify(data, flag)
		return
	}
	cmd := proto.MsgTypeCmd(gotcp.GetCmd(data))
	if handler, ok := this.Cmds[uint64(cmd)]; ok {
		handler(data, flag)
	} else {
		if this.DefaultCmd != nil {
			this.DefaultCmd(data, flag)
		}
	}
}

func (this *SessionIntranet) doVerify(data []byte, flag byte) {
	msg := &proto.MsgVerifyS{}
	if gotcp.DecodeCmd(data, flag, msg) == nil {
		xlog.Errorln("decodeMsg fail.")
		this.Close()
		return
	}
	if msg.GetToken() != xargs.Common.IntranetToken {
		xlog.Errorln("token error.")
		this.Close()
	}
	this.Id = msg.GetId()
	xnodes.Store(this.Id, this)
	this.Verify()
	xlog.Debugln("Id:", msg.GetId(), "verify success.")

	msg.Reset()
	msg.Id = discovery.GetNode().Id()
	msg.Token = xargs.Common.IntranetToken
	this.SendMsg(uint64(proto.MsgTypeCmd_Verify), msg)
}

func (this *SessionIntranet) OnClose() {
	if _, loaded := xnodes.Load(this.Id); loaded {
		xnodes.Delete(this.Id)
	}
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
