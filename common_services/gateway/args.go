package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fananchong/go-x/common"
	"github.com/fananchong/go-x/common/discovery"
	"github.com/fananchong/gotcp"
)

type Args struct {
	common.ArgsBase
	Gateway ArgsGateway
}

type ArgsGateway struct {
	ExternalIp string
	IntranetIp string
	Connect    []int
}

func (this *Args) OnInit() {
	this.Pending.NodeType = int(common.Gateway)                // 设置本节点类型。不为0，则上报自身节点信息到discovery。
	this.Pending.WatchNodeTypes = append(this.Gateway.Connect) // 监视服务节点类型
	this.Pending.ExternalIp = this.Gateway.ExternalIp          // 对外地址
	this.Pending.IntranetIp = this.Gateway.IntranetIp          // 对内地址
	this.adjustExternalPort()                                  // 矫正对外端口
	xnode.SetBaseInfoType(uint32(common.Gateway))              // 设置本节点类型
	xnode.InitPolicy(discovery.Random)                         // 设置选取服务节点策略
}

func NewArgs() *Args {
	return &Args{}
}

func (this *Args) GetDerived() common.IArgs {
	return this
}

func (this *Args) adjustExternalPort() {
	addrinfo := strings.Split(this.Gateway.ExternalIp, ":")
	var port int
	var err error
	if len(addrinfo) < 2 {
		port = gotcp.GetVaildPort()
	} else {
		port, err = strconv.Atoi(addrinfo[1])
		if err != nil {
			panic(err)
			return
		}
	}
	addr := fmt.Sprintf("%s:%d", addrinfo[0], port)
	this.Gateway.ExternalIp = addr
	this.Pending.ExternalIp = addr
}
