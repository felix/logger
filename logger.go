package logger

type Logger interface {
	Log(level Level, args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})

	WithFields(args ...interface{}) Logger
	Named(name string) Logger
}