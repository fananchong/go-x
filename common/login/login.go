package login

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fananchong/go-x/common"
)

type Login struct {
	common.WebService
}

func (this *Login) Start(addr string) {
	this.InitHander()
	this.ListenAndServe(addr)
}

func (this *Login) InitHander() {
	this.HandleFunc("/msg", this.Request)
}

const (
	SIGN1 = "5UY6$f$h"
	SIGN2 = "3wokZB%q"
	SIGN3 = "%2Fi9TRf"
)

func (this *Login) Request(w http.ResponseWriter, req *http.Request) {
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

	s1 := []byte(SIGN1 + paramc[0] + SIGN2 + paramt[0] + SIGN3 + "0.0.1") // TODO: 配置参数待优化
	s2 := md5.Sum(s1)
	s3 := fmt.Sprintf("%x", s2)

	if s3 != params[0] {
		common.GetLogger().Debugln("version error!")
		return
	}
	common.GetLogger().Debugln(c)
	common.GetLogger().Debugln(paramd[0])
}
