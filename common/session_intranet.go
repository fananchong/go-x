package common

import (
	"sync"

	"github.com/fananchong/go-x/common/discovery"
	"github.com/fananchong/go-x/common_services/proto"
	"github.com/fananchong/gotcp"
	proto1 "github.com/golang/protobuf/proto"
)

type SessionIntranet struct {
	gotcp.Session
	Id         string
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
	msg := &proto.MsgVerify{}
	if gotcp.DecodeCmd(data, flag, msg) == nil {
		xlog.Errorln("decodeMsg fail.")
		this.Close()
		return
	}
	if msg.GetToken() != xargs.Common.IntranetToken {
		xlog.Errorln("token error.")
		this.Close()
	}
	this.Id = msg.GetAccount()
	xnodes.Store(this.Id, this)
	this.Verify()
	xlog.Debugln("Id:", msg.GetAccount(), "verify success.")

	msg.Reset()
	msg.Account = discovery.GetNode().Id()
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
		if key.(string) != this.Id {
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
		if key.(string) != this.Id {
			val.(*SessionIntranet).SendMsg(cmd, msg)
		}
		return true
	})
}

var xnodes sync.Map
