package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common_services/proto"
	"github.com/fananchong/gotcp"
	proto1 "github.com/golang/protobuf/proto"
)

func (this *Login) request(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	paramc, ok1 := req.Form["c"]
	paramt, ok2 := req.Form["t"]
	paramd, ok3 := req.Form["d"]
	params, ok4 := req.Form["s"]
	if !ok1 || !ok2 || !ok3 || !ok4 {
		common.GetLogger().Debugln("http request param error!")
		return
	}
	c, err := strconv.Atoi(paramc[0])
	if err != nil {
		common.GetLogger().Debugln("http request param c error!")
		return
	}

	s1 := []byte(xargs.Login.Sign1 + paramc[0] + xargs.Login.Sign2 + paramt[0] + xargs.Login.Sign3 + xargs.Common.Version)
	s2 := md5.Sum(s1)
	s3 := fmt.Sprintf("%x", s2)

	if s3 != params[0] {
		common.GetLogger().Debugln("version error!")
		common.GetLogger().Debugln("   client sign =", params[0])
		common.GetLogger().Debugln("   server sign =", s3)
		common.GetLogger().Debugln("   sign1 =", xargs.Login.Sign1)
		common.GetLogger().Debugln("   sign2 =", xargs.Login.Sign2)
		common.GetLogger().Debugln("   sign3 =", xargs.Login.Sign3)
		common.GetLogger().Debugln("   c =", paramc[0])
		common.GetLogger().Debugln("   t =", paramt[0])
		common.GetLogger().Debugln("   version =", xargs.Common.Version)

		return
	}

	if handler, ok := this.cmds[proto.MsgTypeCmd(c)]; ok {
		handler(w, req, paramd[0], params[0])
	} else {
		common.GetLogger().Debugln("unknow cmd, cmd =", c)
		return
	}
}

func (this *Login) decodeMsg(data string, msg proto1.Message) proto1.Message {
	return gotcp.DecodeCmdEx([]byte(data), 0, msg, 0)
}
