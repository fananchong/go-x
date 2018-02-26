package main

import (
	"github.com/fananchong/go-x/common"
	"github.com/fananchong/gochart"
)

var (
	xlog common.ILogger = common.NewDefaultLogger()
)

type App struct {
	common.App
}

func NewApp() *App {
	this := &App{}
	this.Derived = this
	return this
}

var (
	g_chart *Chart = nil
)

func (this *App) OnAppReady() {
	g_chart = NewChart()
	s := &gochart.ChartServer{}
	s.AddChart("chart", g_chart, false)
	go func() { println(s.ListenAndServe(":8000").Error()) }()

	TcpClient()
	UdpClient()
}

func (this *App) OnAppShutDown() {

}
