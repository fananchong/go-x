package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fananchong/zap"
)

var (
	xlog ILogger = nil
)

type ILogger interface {
	Debug(args ...interface{})
	Debugln(args ...interface{})
	Debugf(format string, args ...interface{})
	Print(args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})
	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(format string, args ...interface{})
	Warning(args ...interface{})
	Warningln(args ...interface{})
	Warningf(format string, args ...interface{})
	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(format string, args ...interface{})
	Init() bool
	Flush()
}

func SetLogger(log ILogger) {
	xlog = log
}

func GetLogger() ILogger {
	return xlog
}

type ZapLogger struct {
	*zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
	return &ZapLogger{}
}

func (this *ZapLogger) Init() bool {
	rawJSON, err := ioutil.ReadFile(GetArgs().Common.LogCfg)
	if err != nil {
		fmt.Println("open logcfg file fail. file:", GetArgs().Common.LogCfg)
		fmt.Println(err)
		return false
	}
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		fmt.Println(err)
		return false
	}
	logger, err := cfg.Build()
	if err != nil {
		fmt.Println(err)
		return false
	}
	this.SugaredLogger = logger.Sugar()
	return true
}

func (this *ZapLogger) Flush() {
	this.SugaredLogger.Sync()
}

type DefaultLogger struct {
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

func (this *DefaultLogger) Debug(args ...interface{}) {
	fmt.Print(args...)
}

func (this *DefaultLogger) Debugln(args ...interface{}) {
	fmt.Println(args...)
}

func (this *DefaultLogger) Debugf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (this *DefaultLogger) Print(args ...interface{}) {
	fmt.Print(args...)
}

func (this *DefaultLogger) Println(args ...interface{}) {
	fmt.Println(args...)
}

func (this *DefaultLogger) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (this *DefaultLogger) Info(args ...interface{}) {
	fmt.Print(args...)
}

func (this *DefaultLogger) Infoln(args ...interface{}) {
	fmt.Println(args...)
}

func (this *DefaultLogger) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (this *DefaultLogger) Warning(args ...interface{}) {
	fmt.Print(args...)
}

func (this *DefaultLogger) Warningln(args ...interface{}) {
	fmt.Println(args...)
}

func (this *DefaultLogger) Warningf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (this *DefaultLogger) Error(args ...interface{}) {
	fmt.Print(args...)
}

func (this *DefaultLogger) Errorln(args ...interface{}) {
	fmt.Println(args...)
}

func (this *DefaultLogger) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (this *DefaultLogger) Fatal(args ...interface{}) {
	fmt.Print(args...)
}

func (this *DefaultLogger) Fatalln(args ...interface{}) {
	fmt.Println(args...)
}

func (this *DefaultLogger) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	panic("")
}

func (this *DefaultLogger) Init() bool {
	return true
}

func (this *DefaultLogger) Flush() {
}
