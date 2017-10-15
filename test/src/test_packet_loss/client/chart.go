package main

import (
	"github.com/fananchong/gochart"
	"strconv"
	"sync"
)

type Chart struct {
	gochart.ChartTime
	tcp []int64
	udp []int64
	m   sync.Mutex
}

func NewChart() *Chart {
	this := &Chart{tcp: make([]int64, DEFAULT_SAMPLE_NUM), udp: make([]int64, DEFAULT_SAMPLE_NUM)}
	this.RefreshTime = strconv.Itoa(DEFAULT_REFRESH_TIME)
	this.ChartType = "line"
	this.Title = "网络丢包测试"
	this.SubTitle = "服务器每100ms发送hello消息给客户端"
	this.YAxisText = "delay"
	this.YMax = "2000"
	this.ValueSuffix = "ms"
	this.PlotLinesY = "{ color:'red', dashStyle:'longdashdot', value:100, width:1, label:{ text:'100ms', align:'left' } }"
	return this
}

func (this *Chart) Update(now int64) []interface{} {

	tcp := make([]int64, DEFAULT_SAMPLE_NUM)
	udp := make([]int64, DEFAULT_SAMPLE_NUM)

	this.m.Lock()
	copy(tcp, this.tcp)
	copy(udp, this.udp)
	this.m.Unlock()

	datas := make([]interface{}, 0)
	json1 := this.AddData("tcp", tcp, now, DEFAULT_SAMPLE_NUM, DEFAULT_REFRESH_TIME)
	datas = append(datas, json1)
	json2 := this.AddData("udp", udp, now, DEFAULT_SAMPLE_NUM, DEFAULT_REFRESH_TIME)
	datas = append(datas, json2)
	return datas
}

func (this *Chart) AddTcpData(v int64) {
	this.m.Lock()
	this.tcp = append(this.tcp, v)
	for len(this.tcp) > DEFAULT_SAMPLE_NUM {
		this.tcp = this.tcp[1:]
	}
	this.m.Unlock()
}

func (this *Chart) AddUdpData(v int64) {
	this.m.Lock()
	this.udp = append(this.udp, v)
	for len(this.udp) > DEFAULT_SAMPLE_NUM {
		this.udp = this.udp[1:]
	}
	this.m.Unlock()
}
