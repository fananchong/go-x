package gochart

type ILogger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
}

type DefaultLogger struct {
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

func (this *DefaultLogger) Error(args ...interface{}) {
}

func (this *DefaultLogger) Errorf(format string, args ...interface{}) {
}

func (this *DefaultLogger) Errorln(args ...interface{}) {
}

func (this *DefaultLogger) Info(args ...interface{}) {
}

func (this *DefaultLogger) Infof(format string, args ...interface{}) {
}

func (this *DefaultLogger) Infoln(args ...interface{}) {
}

var (
	xlog ILogger = NewDefaultLogger()
)

func SetLogger(log ILogger) {
	xlog = log
}
