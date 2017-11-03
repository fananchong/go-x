package main

import (
	"github.com/fananchong/gochart"
	"sync"
)

type Chart struct {
	gochart.ChartTime
	tcp []int64
	udp []int64
	m   sync.Mutex
}

func NewChart() *Chart {
	this := &Chart{tcp: make([]int64, 0), udp: make([]int64, 0)}
	this.TickUnit = 100
	this.RefreshTime = DEFAULT_REFRESH_TIME
	this.SampleNum = DEFAULT_SAMPLE_NUM
	this.ChartType = "line"
	this.Title = "网络丢包测试"
	this.SubTitle = "服务器每100ms发送hello消息给客户端"
	this.YAxisText = "delay"
	this.YMax = "2000"
	this.ValueSuffix = "ms"
	this.TickLabelStep = "100"
	this.PlotLinesY = "{ color:'red', dashStyle:'longdashdot', value:100, width:1, label:{ text:'100ms', align:'left' } }"
	this.PlotLinesY += ",{ color:'red', dashStyle:'longdashdot', value:200, width:1, label:{ text:'200ms', align:'left' } }"
	return this
}

func (this *Chart) Update(now int64) map[string][]interface{} {
	datas := make(map[string][]interface{})
	this.m.Lock()
	datas["tcp"] = make([]interface{}, 0)
	for _, v := range this.tcp {
		datas["tcp"] = append(datas["tcp"], v)
	}
	datas["udp"] = make([]interface{}, 0)
	for _, v := range this.udp {
		datas["udp"] = append(datas["udp"], v)
	}
	this.tcp = this.tcp[:0]
	this.udp = this.udp[:0]
	this.m.Unlock()
	return datas
}

func (this *Chart) AddTcpData(v int64) {
	this.m.Lock()
	this.tcp = append(this.tcp, v)
	this.m.Unlock()
}

func (this *Chart) AddUdpData(v int64) {
	this.m.Lock()
	this.udp = append(this.udp, v)
	this.m.Unlock()
}
