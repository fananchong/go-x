package common

import (
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
	Init()
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
	log *zap.Logger
}

func NewZapLogger() *ZapLogger {
	return &ZapLogger{}
}

func (this *ZapLogger) Init() {
}

func (this *ZapLogger) Flush() {
}
