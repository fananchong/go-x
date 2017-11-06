package common

import (
	gloglog "github.com/golang/glog"
)

type ILogger interface {
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
	Flush()
}

type DefaultLogger struct {
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

func (this *DefaultLogger) Info(args ...interface{}) {
}

func (this *DefaultLogger) Infof(format string, args ...interface{}) {
}

func (this *DefaultLogger) Infoln(args ...interface{}) {
}

func (this *DefaultLogger) Warning(args ...interface{}) {
}

func (this *DefaultLogger) Warningln(args ...interface{}) {
}

func (this *DefaultLogger) Warningf(format string, args ...interface{}) {
}

func (this *DefaultLogger) Error(args ...interface{}) {
}

func (this *DefaultLogger) Errorf(format string, args ...interface{}) {
}

func (this *DefaultLogger) Errorln(args ...interface{}) {
}

func (this *DefaultLogger) Fatal(args ...interface{}) {
}

func (this *DefaultLogger) Fatalln(args ...interface{}) {
}

func (this *DefaultLogger) Fatalf(format string, args ...interface{}) {
}

func (this *DefaultLogger) Flush() {
}

type GLog struct {
}

func NewGLog() *GLog {
	return &GLog{}
}

func (this *GLog) Info(args ...interface{}) {
	gloglog.Info(args)
}

func (this *GLog) Infof(format string, args ...interface{}) {
	gloglog.Infof(format, args)
}

func (this *GLog) Infoln(args ...interface{}) {
	gloglog.Infoln(args)
}

func (this *GLog) Warning(args ...interface{}) {
	gloglog.Warning(args)
}

func (this *GLog) Warningln(args ...interface{}) {
	gloglog.Warningln(args)
}

func (this *GLog) Warningf(format string, args ...interface{}) {
	gloglog.Warningf(format, args)
}

func (this *GLog) Error(args ...interface{}) {
	gloglog.Error(args)
}

func (this *GLog) Errorf(format string, args ...interface{}) {
	gloglog.Errorf(format, args)
}

func (this *GLog) Errorln(args ...interface{}) {
	gloglog.Errorln(args)
}

func (this *GLog) Fatal(args ...interface{}) {
	gloglog.Fatal(args)
}

func (this *GLog) Fatalln(args ...interface{}) {
	gloglog.Fatalln(args)
}

func (this *GLog) Fatalf(format string, args ...interface{}) {
	gloglog.Fatalf(format, args)
}

func (this *GLog) Flush() {
	gloglog.Flush()
}

var (
	glog ILogger = NewDefaultLogger()
)

func SetLogger(log ILogger) {
	glog = log
}
