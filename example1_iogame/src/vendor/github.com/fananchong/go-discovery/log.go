package godiscovery

type ILogger interface {
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
