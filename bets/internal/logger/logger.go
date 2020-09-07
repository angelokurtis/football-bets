package logger

type Logger interface {
	Info(msg string)
	Infof(format string, v ...interface{})
	Debug(msg string)
	Debugf(format string, v ...interface{})
	Error(err error)
	Fatal(err error)
}
