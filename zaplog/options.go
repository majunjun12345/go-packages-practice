package zaplog

import "go.uber.org/zap/zapcore"

// OffStdout 关闭控制台输出
func OffStdout() ModOptions {
	return func(option *options) {
		option.isStdout = false
	}
}

// SetMaxSize 设置单个日志文件的大小
func SetMaxSize(MaxSize int) ModOptions {
	return func(option *options) {
		option.archive.maxSize = MaxSize
	}
}

// SetMaxBackups 设置日志文件备份数
func SetMaxBackups(MaxBackups int) ModOptions {
	return func(option *options) {
		option.archive.maxBackups = MaxBackups
	}
}

// SetMaxAge 设置日志保存时限
func SetMaxAge(MaxAge int) ModOptions {
	return func(option *options) {
		option.archive.maxAge = MaxAge
	}
}

// SetCompress 是否压缩备份的日志文件
func SetCompress() ModOptions {
	return func(option *options) {
		option.archive.compress = true
	}
}

// SetLogFileDir 设置日志文件保存目录
func SetLogFileDir(LogFileDir string) ModOptions {
	return func(option *options) {
		option.archive.logFileDir = LogFileDir
	}
}

// SetLevel 设置日志级别
func SetLevel(Level zapcore.Level) ModOptions {
	return func(option *options) {
		option.level = Level
	}
}

// SetErrorFileName 设置 error 级别日志的写入文件
func SetErrorFileName(ErrorFileName string) ModOptions {
	return func(option *options) {
		option.archive.errorFileName = ErrorFileName
	}
}

// SetWarnFileName 设置 warn 级别日志的写入文件
func SetWarnFileName(WarnFileName string) ModOptions {
	return func(option *options) {
		option.archive.warnFileName = WarnFileName
	}
}

// SetInfoFileName 设置 info 级别日志的写入文件
func SetInfoFileName(InfoFileName string) ModOptions {
	return func(option *options) {
		option.archive.infoFileName = InfoFileName
	}
}

// SetDebugFileName 设置 debug 级别日志的写入文件
func SetDebugFileName(DebugFileName string) ModOptions {
	return func(option *options) {
		option.archive.debugFileName = DebugFileName
	}
}

// SetDevelopment 设置是否是开发模式
func SetDevelopment(Development bool) ModOptions {
	return func(option *options) {
		option.development = Development
	}
}

// SetInitialFields 设置日志的默认字段
func SetInitialFields(key string, value interface{}) ModOptions {
	return func(option *options) {
		if option.initialFields == nil {
			option.initialFields = make(map[string]interface{})
		}
		option.initialFields[key] = value
	}
}

// SetMessageKey set EncoderConfig MessageKey
func SetMessageKey(key string) ModOptions {
	return func(option *options) {
		option.messageKey = key
	}
}

// SetLevelKey set EncoderConfig LevelKey
func SetLevelKey(key string) ModOptions {
	return func(option *options) {
		option.levelKey = key
	}
}

// SetTimeKey set EncoderConfig TimeKey
func SetTimeKey(key string) ModOptions {
	return func(option *options) {
		option.timeKey = key
	}
}

// SetNameKey set EncoderConfig NameKey
func SetNameKey(key string) ModOptions {
	return func(option *options) {
		option.nameKey = key
	}
}

// SetCallerKey set EncoderConfig CallerKey
func SetCallerKey(key string) ModOptions {
	return func(option *options) {
		option.callerKey = key
	}
}

// SetStacktraceKey set EncoderConfig StacktraceKey
func SetStacktraceKey(key string) ModOptions {
	return func(option *options) {
		option.stacktraceKey = key
	}
}

// SetLineEnding set EncoderConfig LineEnding
func SetLineEnding(lineEnding string) ModOptions {
	return func(option *options) {
		option.lineEnding = lineEnding
	}
}

// SetNoStacktrace 不打印堆栈信息
func SetNoStacktrace(b bool) ModOptions {
	return func(option *options) {
		option.disableStacktrace = b
	}
}
