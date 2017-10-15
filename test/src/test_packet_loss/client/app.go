package main

import (
	"common"
	"github.com/fananchong/gochart"
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
	s.AddChart("chart", g_chart)
	go func() { println(s.ListenAndServe(":8000").Error()) }()

	go TcpClient()
	go UdpClient()
}

func (this *App) OnAppShutDown() {

}
