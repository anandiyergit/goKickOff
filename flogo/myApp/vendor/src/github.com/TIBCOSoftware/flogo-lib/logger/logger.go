package logger

import "errors"

type Logger interface {
    Debug(args ...interface{})
    Debugf(format string, args ...interface{})
    Info(args ...interface{})
    Infof(format string, args ...interface{})
    Warn(args ...interface{})
    Warnf(format string, args ...interface{})
    Error(args ...interface{})
    Errorf(format string, args ...interface{})
    SetLogLevel(Level)
}

type LoggerFactory interface {
	GetLogger(name string) (Logger, error)
}

type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
)

var logFactory LoggerFactory

func RegisterLoggerFactory(factory LoggerFactory) {
	logFactory = factory
}

func GetLogger(name string) (Logger, error) {
	if logFactory == nil {
		return nil, errors.New("No logger factory found.")
	}
	return logFactory.GetLogger(name)
}
