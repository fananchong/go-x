package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fananchong/go-x/base"
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
		base.XLOG.Debugln("http request param error!")
		return
	}
	c, err := strconv.Atoi(paramc[0])
	if err != nil {
		base.XLOG.Debugln("http request param c error!")
		return
	}

	s1 := []byte(externArgs.Sign1 + paramc[0] + externArgs.Sign2 + paramt[0] + externArgs.Sign3 + base.XARGS.Common.Version)
	s2 := md5.Sum(s1)
	s3 := fmt.Sprintf("%x", s2)

	if s3 != params[0] {
		base.XLOG.Debugln("version error!")
		base.XLOG.Debugln("   client sign =", params[0])
		base.XLOG.Debugln("   server sign =", s3)
		base.XLOG.Debugln("   sign1 =", externArgs.Sign1)
		base.XLOG.Debugln("   sign2 =", externArgs.Sign2)
		base.XLOG.Debugln("   sign3 =", externArgs.Sign3)
		base.XLOG.Debugln("   c =", paramc[0])
		base.XLOG.Debugln("   t =", paramt[0])
		base.XLOG.Debugln("   version =", base.XARGS.Common.Version)

		return
	}

	if handler, ok := this.cmds[proto.MsgTypeCmd(c)]; ok {
		handler(w, req, paramd[0], params[0])
	} else {
		base.XLOG.Debugln("unknow cmd, cmd =", c)
		return
	}
}

func (this *Login) decodeMsg(data string, msg proto1.Message) proto1.Message {
	return gotcp.DecodeCmdEx([]byte(data), 0, msg, 0)
}
