package common

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"syscall"
	"time"

	"github.com/fananchong/go-x/base"
	godiscovery "github.com/fananchong/go-x/internal/common/k8s"
	discovery "github.com/fananchong/go-x/internal/common/k8s/serverlist"
	"github.com/fatih/structs"
)

type AppInterface interface {
	OnAppReady()
	OnAppShutDown()
}

type App struct {
	signal  chan os.Signal
	Derived AppInterface
	Type    int
	Node    interface{}
}

func (this *App) Run(argsObj interface{}) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GC()

	termination := false
	this.signal = make(chan os.Signal)
	signal.Notify(this.signal, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGPIPE)

	if base.XLOG == nil {
		base.XLOG = base.NewDefaultLogger()
	}
	defer base.XLOG.Flush()

	this.initArgs(argsObj)
	this.initProf()
	this.initNode()

	this.Derived.OnAppReady()

	for !termination {
		select {
		case sig := <-this.signal:
			switch sig {
			case syscall.SIGPIPE:
			default:
				termination = true
			}
			base.XLOG.Infoln("[app] recive signal. signal no =", sig)
		}
	}

	this.Derived.OnAppShutDown()
	this.closeDetail()
}

func (this *App) Close() {
	close(this.signal)
}

func (this *App) closeDetail() {
	base.XLOG.Infoln("app close ...")
	node := this.Node.(godiscovery.INode).GetBase().(*discovery.Node)
	node.Close()
}

func (this *App) initLog() {
	base.XLOG = base.NewGLogger()
	if base.XARGS.Common.LogDir != "" {
		os.MkdirAll(base.XARGS.Common.LogDir, os.ModeDir)
	}
	base.XLOG.SetLogDir(base.XARGS.Common.LogDir)
	base.XLOG.SetLogLevel(base.XARGS.Common.LogLevel)
}

func (this *App) initArgs(argsObj interface{}) {
	if argsObj == nil {
		return
	}
	base.XARGS = argsObj.(*base.ArgsBase)
	parseArgs(argsObj)
	this.initLog()

	fields := structs.Fields(argsObj)
	this.initArgsDetail(fields)

	// OnInit
	f := reflect.ValueOf(argsObj).MethodByName("OnInit")
	if f.IsValid() {
		f.Call([]reflect.Value{})
	}
}

func (this *App) initArgsDetail(fields []*structs.Field) {
	for _, field := range fields {
		if field.Name() == "ArgsBase" {
			continue
		}
		switch field.Kind() {
		case reflect.Struct:
			fields2 := structs.Fields(field.Value())
			this.initArgsDetail(fields2)
		default:
			switch field.Name() {
			case "Connect":
				base.XARGS.Pending.WatchNodeTypes = append(base.XARGS.Pending.WatchNodeTypes, field.Value().([]int)...)
			}
		}
	}
}

func (this *App) initNode() {
	if base.XARGS != nil {
		initServerType()
		args := base.XARGS
		if this.Type == 0 &&
			len(args.Pending.WatchNodeTypes) == 0 {
			return
		}
		if this.Node == nil {
			this.Node = discovery.NewNode()
		}
		node := this.Node.(godiscovery.INode).GetBase().(*discovery.Node)
		node.Init(this.Type, args.Pending.WatchNodeTypes, 5*time.Second, this.Node.(godiscovery.INode))
		discovery.SetNode(node)
	}
}

func (this *App) initProf() {
	if base.XARGS != nil && base.XARGS.Common.Debug {
		port := 58000 + this.Type
		go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}
}
