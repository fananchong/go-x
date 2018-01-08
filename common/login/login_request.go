package login

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strconv"

	pb "github.com/fananchong/go-proto-helper"
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/proto"
	proto1 "github.com/golang/protobuf/proto"
)

const (
	sign1 = "5UY6$f$h"
	sign2 = "3wokZB%q"
	sign3 = "%2Fi9TRf"
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

	s1 := []byte(sign1 + paramc[0] + sign2 + paramt[0] + sign3 + "0.0.1") // TODO: 配置参数待优化
	s2 := md5.Sum(s1)
	s3 := fmt.Sprintf("%x", s2)

	if s3 != params[0] {
		common.GetLogger().Debugln("version error!")
		return
	}

	if handler, ok := this.cmds[proto.MsgTypeCmd(c)]; ok {
		handler(w, req, paramd[0])
	} else {
		common.GetLogger().Debugln("unknow cmd, cmd =", c)
		return
	}
}

func (this *Login) decodeMsg(data string, msg proto1.Message) proto1.Message {
	return pb.DecodeCmdEx([]byte(data), 0, &proto.MsgLogin{}, 0)
}
