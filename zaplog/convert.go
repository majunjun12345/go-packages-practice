package zaplog

import "go.uber.org/zap/zapcore"

type Level string

// Unmarshal 将配置文件中的level定义转换为 zapcore Level
func (l Level) Unmarshal(level string) zapcore.Level {
	switch level {
	case "debug", "Debug", "DEBUG":
		return zapcore.DebugLevel
	case "info", "Info", "INFO":
		return zapcore.InfoLevel
	case "warn", "Warn", "WARN":
		return zapcore.WarnLevel
	case "error", "Error", "ERROR":
		return zapcore.ErrorLevel
	case "panic", "Panic", "PANIC":
		return zapcore.PanicLevel
	case "fatal", "Fatal", "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

type Env string

// Unmarshal 将配置文件中定义的 env 转换为 zapConfig.Development 的 true 或 false
func (e Env) Unmarshal(env string) bool {
	switch env {
	case "dev", "Dev", "DEV":
		return true
	case "pro", "Pro", "PRO", "product", "Product", "PRODUCT":
		return false
	default:
		return true
	}
}
