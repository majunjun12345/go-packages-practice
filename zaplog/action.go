package zaplog

import "go.uber.org/zap"

// 对外暴露的六个 log 方法

// Debug will log at debug level
func Debug(msg string, fields ...zap.Field) {
	globalLoggerWarp().Debug(msg, fields...)
}

// Info logs a message at InfoLevel
func Info(msg string, fields ...zap.Field) {
	globalLoggerWarp().Info(msg, fields...)
}

// Warn will log stacktrace info
func Warn(msg string, fields ...zap.Field) {
	globalLoggerWarp().Warn(msg, fields...)
}

// Error will log stacktrace info
func Error(msg string, fields ...zap.Field) {
	globalLoggerWarp().Error(msg, fields...)
}

// Panic The logger then panics, even if logging at PanicLevel is disabled, will recovery if set
func Panic(msg string, fields ...zap.Field) {
	globalLoggerWarp().Panic(msg, fields...)
}

// Fatal The logger then calls os.Exit(1)
func Fatal(msg string, fields ...zap.Field) {
	globalLoggerWarp().Fatal(msg, fields...)
}

func globalLoggerWarp() *zap.Logger {
	return globalLogger.Logger.WithOptions(zap.AddCallerSkip(1))
}
