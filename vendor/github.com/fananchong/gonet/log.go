package gonet

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
}

type DefaultLogger struct {
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
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

var (
	xlog ILogger = NewDefaultLogger()
)

func SetLogger(log ILogger) {
	xlog = log
}
