package glog

import "fmt"

func (l *loggingT) Init() bool {
	return true
}

func (l *loggingT) Flush() {
	l.lockAndFlushAll()
}

func (l *loggingT) Debug(args ...interface{}) {
	fmt.Print(args...)
}

func (l *loggingT) Debugln(args ...interface{}) {
	fmt.Println(args...)
}

func (l *loggingT) Debugf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *loggingT) Print(args ...interface{}) {
	l.print(infoLog, args...)
}

func (l *loggingT) Println(args ...interface{}) {
	l.println(infoLog, args...)
}

func (l *loggingT) Printf(format string, args ...interface{}) {
	l.printf(infoLog, format, args...)
}

func (l *loggingT) Info(args ...interface{}) {
	l.print(infoLog, args...)
}

func (l *loggingT) Infoln(args ...interface{}) {
	l.println(infoLog, args...)
}

func (l *loggingT) Infof(format string, args ...interface{}) {
	l.printf(infoLog, format, args...)
}

func (l *loggingT) Warning(args ...interface{}) {
	l.print(warningLog, args...)
}

func (l *loggingT) Warningln(args ...interface{}) {
	l.println(warningLog, args...)
}

func (l *loggingT) Warningf(format string, args ...interface{}) {
	l.printf(warningLog, format, args...)
}

func (l *loggingT) Error(args ...interface{}) {
	l.print(errorLog, args...)
}

func (l *loggingT) Errorln(args ...interface{}) {
	l.println(errorLog, args...)
}

func (l *loggingT) Errorf(format string, args ...interface{}) {
	l.printf(errorLog, format, args...)
}

func (l *loggingT) Fatal(args ...interface{}) {
	l.print(fatalLog, args...)
}

func (l *loggingT) Fatalln(args ...interface{}) {
	l.println(fatalLog, args...)
}

func (l *loggingT) Fatalf(format string, args ...interface{}) {
	l.printf(fatalLog, format, args...)
}

func (l *loggingT) SetLogLevel(level int) {
	l.stderrThreshold = severity(level)
}

func (l *loggingT) SetLogDir(dir *string) {
	logDir = dir
}

type LoggingT = loggingT

func GetLogger() *LoggingT {
	return &logging
}
