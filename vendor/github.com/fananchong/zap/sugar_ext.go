package zap

func (s *SugaredLogger) Debugln(args ...interface{}) {
	s.log(DebugLevel, "", args, nil)
}

func (s *SugaredLogger) Print(args ...interface{}) {
	s.log(InfoLevel, "", args, nil)
}

func (s *SugaredLogger) Println(args ...interface{}) {
	s.log(InfoLevel, "", args, nil)
}

func (s *SugaredLogger) Printf(template string, args ...interface{}) {
	s.log(InfoLevel, template, args, nil)
}

func (s *SugaredLogger) Infoln(args ...interface{}) {
	s.log(InfoLevel, "", args, nil)
}

func (s *SugaredLogger) Warning(args ...interface{}) {
	s.log(WarnLevel, "", args, nil)
}

func (s *SugaredLogger) Warningln(args ...interface{}) {
	s.log(WarnLevel, "", args, nil)
}

func (s *SugaredLogger) Warningf(template string, args ...interface{}) {
	s.log(WarnLevel, template, args, nil)
}

func (s *SugaredLogger) Errorln(args ...interface{}) {
	s.log(ErrorLevel, "", args, nil)
}

func (s *SugaredLogger) Fatalln(args ...interface{}) {
	s.log(FatalLevel, "", args, nil)
}
