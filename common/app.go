package common

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	godiscovery "github.com/fananchong/go-discovery"
)

type AppInterface interface {
	OnAppReady()
	OnAppShutDown()
}

type App struct {
	Derived AppInterface
	Args    IArgs
	Node    interface{}
}

func (this *App) Run() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GC()

	termination := false
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGPIPE)

	this.initArgs()
	this.initNode()

	this.Derived.OnAppReady()

	for !termination {
		select {
		case sig := <-ch:
			switch sig {
			case syscall.SIGPIPE:
			default:
				termination = true
			}
			xlog.Infoln("[app] recive signal. signal no =", sig)
		}
	}

	this.Derived.OnAppShutDown()

	xlog.Flush()
}

func (this *App) initArgs() {
	if this.Args == nil {
		xlog.Warningln("Need New Args Object!")
		return
	}
	this.Args.Init()
	flag.Parse()
	this.Args.Parse()
}

func (this *App) initNode() {
	if this.Node != nil && this.Args != nil {
		node := this.Node.(godiscovery.INode)
		node.Init(this.Node)
		node.SetLogger(xlog)
		args := this.Args.GetBase()
		node.OpenByStr(args.EtcdHosts, int(args.EtcdNodeType), args.EtcdWatchNodeTypes, args.EtcdPutInterval)
	} else {
		xlog.Errorln("Need New Args Object OR Node Object!")
	}
}
