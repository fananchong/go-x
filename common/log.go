package common

import (
	"fmt"

	"github.com/golang/glog"
)

var (
	xlog ILogger = NewGLog()
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
	Flush()
}

type DefaultLogger struct {
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

func (this *DefaultLogger) Debug(args ...interface{}) {
}

func (this *DefaultLogger) Debugln(args ...interface{}) {
}

func (this *DefaultLogger) Debugf(format string, args ...interface{}) {
}

func (this *DefaultLogger) Print(args ...interface{}) {
}

func (this *DefaultLogger) Println(args ...interface{}) {
}

func (this *DefaultLogger) Printf(format string, args ...interface{}) {
}

func (this *DefaultLogger) Info(args ...interface{}) {
}

func (this *DefaultLogger) Infoln(args ...interface{}) {
}

func (this *DefaultLogger) Infof(format string, args ...interface{}) {
}

func (this *DefaultLogger) Warning(args ...interface{}) {
}

func (this *DefaultLogger) Warningln(args ...interface{}) {
}

func (this *DefaultLogger) Warningf(format string, args ...interface{}) {
}

func (this *DefaultLogger) Error(args ...interface{}) {
}

func (this *DefaultLogger) Errorln(args ...interface{}) {
}

func (this *DefaultLogger) Errorf(format string, args ...interface{}) {
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

func (this *GLog) Debug(args ...interface{}) {
	fmt.Print(args...)
}

func (this *GLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (this *GLog) Debugln(args ...interface{}) {
	fmt.Println(args...)
}

func (this *GLog) Print(args ...interface{}) {
	glog.Info(args...)
}

func (this *GLog) Printf(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

func (this *GLog) Println(args ...interface{}) {
	glog.Infoln(args...)
}

func (this *GLog) Info(args ...interface{}) {
	glog.Info(args...)
}

func (this *GLog) Infof(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

func (this *GLog) Infoln(args ...interface{}) {
	glog.Infoln(args...)
}

func (this *GLog) Warning(args ...interface{}) {
	glog.Warning(args...)
}

func (this *GLog) Warningln(args ...interface{}) {
	glog.Warningln(args...)
}

func (this *GLog) Warningf(format string, args ...interface{}) {
	glog.Warningf(format, args...)
}

func (this *GLog) Error(args ...interface{}) {
	glog.Error(args...)
}

func (this *GLog) Errorf(format string, args ...interface{}) {
	glog.Errorf(format, args...)
}

func (this *GLog) Errorln(args ...interface{}) {
	glog.Errorln(args...)
}

func (this *GLog) Fatal(args ...interface{}) {
	glog.Fatal(args...)
}

func (this *GLog) Fatalln(args ...interface{}) {
	glog.Fatalln(args...)
}

func (this *GLog) Fatalf(format string, args ...interface{}) {
	glog.Fatalf(format, args...)
}

func (this *GLog) Flush() {
	glog.Flush()
}

func SetLogger(log ILogger) {
	xlog = log
}

func GetLogger() ILogger {
	return xlog
}
