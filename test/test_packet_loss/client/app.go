package main

import "github.com/fananchong/gochart"

var (
	g_chart *Chart = nil
)

func main() {
	g_chart = NewChart()
	s := &gochart.ChartServer{}
	s.AddChart("chart", g_chart, false)
	go func() { println(s.ListenAndServe(":8000").Error()) }()

	TcpClient()
	UdpClient()
}
