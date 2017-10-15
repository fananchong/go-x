package common

import (
	"flag"
	"github.com/golang/glog"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

type AppInterface interface {
	OnAppReady()
	OnAppShutDown()
}

type App struct {
	Derived AppInterface
}

func (this *App) Run() {
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GC()

	termination := false
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGPIPE)

	this.Derived.OnAppReady()

	for !termination {
		select {
		case sig := <-ch:
			switch sig {
			case syscall.SIGPIPE:
			default:
				termination = true
			}
			glog.Infoln("[app] recive signal. signal no =", sig)
		}
	}

	this.Derived.OnAppShutDown()

	glog.Flush()
}
