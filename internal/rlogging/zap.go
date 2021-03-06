package rlogging

import "go.uber.org/zap"

type ZapLogger struct {
	*zap.SugaredLogger
}

func (l *ZapLogger) Infof(tmpl string, args ...interface{}) {
	l.SugaredLogger.Infof(tmpl, args)
}

func (l *ZapLogger) Debugf(tmpl string, args ...interface{}) {
	l.SugaredLogger.Debugf(tmpl, args)
}

func (l *ZapLogger) Errorf(tmpl string, args ...interface{}) {
	l.SugaredLogger.Errorf(tmpl, args)
}
