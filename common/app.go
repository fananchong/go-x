package common

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	godiscovery "github.com/fananchong/go-discovery"
	"github.com/fananchong/go-x/common/discovery"
	"github.com/fananchong/gotcp"
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
	Args    IArgs
	Logger  ILogger
	Node    interface{}
}

func (this *App) Run() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GC()

	termination := false
	this.signal = make(chan os.Signal)
	signal.Notify(this.signal, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGPIPE)

	if this.Logger == nil {
		this.Logger = NewDefaultLogger()
	}
	defer this.Logger.Flush()

	this.initArgs()
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
			xlog.Infoln("[app] recive signal. signal no =", sig)
		}
	}

	this.Derived.OnAppShutDown()
	this.closeDetail()
}

func (this *App) Close() {
	close(this.signal)
}

func (this *App) closeDetail() {
	xlog.Infoln("app close ...")
	node := this.Node.(godiscovery.INode).GetBase().(*discovery.Node)
	node.Close()
}

func (this *App) initLog() {
	if this.Args.GetBase().Common.LogDir != "" {
		os.MkdirAll(this.Args.GetBase().Common.LogDir, os.ModeDir)
	}
	this.Logger.SetLogDir(this.Args.GetBase().Common.LogDir)
	this.Logger.SetLogLevel(this.Args.GetBase().Common.LogLevel)
	SetLogger(this.Logger)
}

func (this *App) initArgs() {
	if this.Args == nil {
		return
	}
	this.Args.Init(this.Args)
	this.initLog()

	fields := structs.Fields(this.Args)
	this.initArgsDetail(fields)

	// OnInit
	f := reflect.ValueOf(this.Args).MethodByName("OnInit")
	if f.IsValid() {
		f.Call([]reflect.Value{})
	}

	SetArgs(this.Args.GetBase())
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
			case "ExternalIp":
				externalIp := field.Value().(string)
				addrinfo := strings.Split(externalIp, ":")
				var port int
				var err error
				if len(addrinfo) < 2 {
					port = gotcp.GetVaildPort()
				} else {
					port, err = strconv.Atoi(addrinfo[1])
					if err != nil {
						panic(err)
						return
					}
				}
				this.Args.GetBase().Pending.ExternalIp = fmt.Sprintf("%s:%d", addrinfo[0], port)
			case "IntranetIp":
				this.Args.GetBase().Pending.IntranetIp = field.Value().(string)
			case "Connect":
				this.Args.GetBase().Pending.WatchNodeTypes = append(this.Args.GetBase().Pending.WatchNodeTypes, field.Value().([]int)...)
			}
		}
	}
	this.Args.GetBase().Pending.NodeType = this.Type
}

func (this *App) initNode() {
	if this.Args != nil {
		args := this.Args.GetBase()
		if args.Pending.NodeType == 0 &&
			len(args.Pending.WatchNodeTypes) == 0 &&
			len(args.Etcd.Hosts) == 0 {
			return
		}
		if this.Node == nil {
			this.Node = &discovery.Node{}
		}
		node := this.Node.(godiscovery.INode).GetBase().(*discovery.Node)
		node.SetBaseInfoType(uint32(this.Type))
		node.InitPolicy(discovery.Ordered)
		node.SetBaseInfoIP(args.Pending.ExternalIp, args.Pending.IntranetIp)
		discovery.SetLogger(xlog)
		node.Init(this.Node)
		node.Open(args.Etcd.Hosts, args.Pending.NodeType, args.Pending.WatchNodeTypes, int64(args.Etcd.PutInterval))
		discovery.SetNode(node)
	}
}

func (this *App) initProf() {
	if this.Args != nil && this.Args.GetBase().Common.Debug {
		port := 58000 + this.Args.GetBase().Pending.NodeType
		go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}
}
