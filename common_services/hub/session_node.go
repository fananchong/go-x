package main

import "github.com/fananchong/gotcp"

type SessionNode struct {
	gotcp.Session
}

func (this *SessionNode) OnRecv(data []byte, flag byte) {
	if this.IsVerified() == false {
		this.Verify()
	}
}

func (this *SessionNode) OnClose() {

}
