package common

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	godiscovery "github.com/fananchong/go-discovery"
	"github.com/fananchong/go-x/common/discovery"
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
	this.initLog()
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

func (this *App) initLog() {
	arg := flag.Lookup("log_dir")
	if arg != nil && arg.Value != nil {
		if arg.Value.String() != "" {
			os.MkdirAll(arg.Value.String(), os.ModeDir)
		}
	}
}

func (this *App) initArgs() {
	if this.Args == nil {
		xlog.Warningln("Need New Args Object!")
		return
	}
	this.Args.Init(this.Args.GetDerived())
}

func (this *App) initNode() {
	if this.Node != nil && this.Args != nil {
		args := this.Args.GetBase()
		node := this.Node.(godiscovery.INode).GetBase().(*discovery.Node)
		node.SetBaseInfoIP(args.ExternalIp, args.IntranetIp)
		node.SetLogger(xlog)
		node.Init(this.Node)
		node.Open(args.Etcd.Hosts, args.Etcd.NodeType, args.Etcd.WatchNodeTypes, int64(args.Etcd.PutInterval))
	} else {
		xlog.Errorln("Need New Args Object OR Node Object!")
	}
}
