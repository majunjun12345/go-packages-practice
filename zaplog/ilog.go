package zaplog

import "fmt"

/*
 type Logger interface {
	Printf(format string, v ...interface{})
}
*/
// Printf 实现 es 的 Logger
func (l *ZapLogger) Printf(format string, v ...interface{}) {
	globalLoggerWarp().Info(fmt.Sprintf(format, v...))
}

// ----------------------------------------------------------------
type Logger interface {
	// Error logs a message at error priority
	Error(msg string)

	// Infof logs a message at info priority
	Infof(msg string, args ...interface{})
}

// 实现 opentracing 的 Loggger
func (l *ZapLogger) Error(msg string) {
	globalLoggerWarp().Info(msg)
}

func (l *ZapLogger) Infof(msg string, args ...interface{}) {
	globalLoggerWarp().Info(fmt.Sprintf(msg, args...))
}
