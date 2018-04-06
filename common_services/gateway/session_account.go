package main

import "github.com/fananchong/gotcp"

type SessionAccount struct {
	gotcp.Session
}

func (this *SessionAccount) OnRecv(data []byte, flag byte) {
	if this.IsVerified() == false {
		this.Verify()
	}
}

func (this *SessionAccount) OnClose() {

}
