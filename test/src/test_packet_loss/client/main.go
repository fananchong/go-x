package main

// usage: client --tcpaddr=192.168.1.3:3000 --udpaddr=192.168.1.3:3000 --log_dir=./log -stderrthreshold 0

import (
	"flag"
)

var (
	MSG                  = []byte("hello")
	DEFAULT_REFRESH_TIME = 1
	DEFAULT_SAMPLE_NUM   = 2 * 60 / DEFAULT_REFRESH_TIME * 10
)

func main() {
	tcpaddr := ""
	flag.StringVar(&tcpaddr, "tcpaddr", "localhost:3000", "tcp server address")
	udpaddr := ""
	flag.StringVar(&udpaddr, "udpaddr", "localhost:3000", "udp server address")

	app := NewApp()
	app.Run()
}
