package login

import (
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
	this.HandleFunc("/login", this.Request)
}

func (this *Login) Request(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	paramcmd, ok1 := req.Form["cmd"]
	paramdata, ok2 := req.Form["data"]
	if !ok1 || !ok2 {
		common.GetLogger().Debugln("http request param error!")
		return
	}
	cmd, err := strconv.Atoi(paramcmd[0])
	if err != nil {
		common.GetLogger().Debugln("http request param cmd error!")
		return
	}
	data := paramdata[0]

	common.GetLogger().Debugln(cmd)
	common.GetLogger().Debugln(data)
}
