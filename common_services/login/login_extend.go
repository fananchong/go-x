package main

import (
	"github.com/fananchong/go-x/common_services/proto"
)

// 第3方平台，获取帐号密码
func (this *Login) GetPassword(account string, mode proto.LoginMode, userdata []byte) (string, proto.LoginError) {
	return "", proto.LoginError_ErrPlatformSide
}
