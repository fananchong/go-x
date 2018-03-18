package gotcp

import "fmt"

type ILogger interface {
	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(format string, args ...interface{})
}

type DefaultLogger struct {
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
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

func (this *DefaultLogger) Error(args ...interface{}) {
	fmt.Print(args...)
}

func (this *DefaultLogger) Errorln(args ...interface{}) {
	fmt.Println(args...)
}

func (this *DefaultLogger) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

var (
	xlog ILogger = NewDefaultLogger()
)

func SetLogger(log ILogger) {
	xlog = log
}
