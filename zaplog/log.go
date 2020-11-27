package zaplog

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger *ZapLogger = &ZapLogger{}
	sp                      = string(filepath.Separator)

	errWS, warnWS, infoWS, debugWS zapcore.WriteSyncer       // 文件IO输出
	debugConsoleWS                 = zapcore.Lock(os.Stdout) // 控制台标准输出
	errorConsoleWS                 = zapcore.Lock(os.Stderr)
)

// ModOptions 日志配置
type ModOptions func(opts *options)

// 日志文件归档配置
type archiveConfig struct {
	logFileDir    string //文件保存目录
	errorFileName string
	warnFileName  string
	infoFileName  string
	debugFileName string
	maxSize       int  // 按大小切割（M）
	maxBackups    int  // 默认备份数
	maxAge        int  // 保存的最大天数
	compress      bool // 是否对日志进行压缩
}

// Options 日志 conf 选项
type options struct {
	isStdout bool // 是否在控制台输出

	archive archiveConfig

	// zap.Config
	level             zapcore.Level          // 日志等级
	development       bool                   // 是否是开发模式, 如果是, 则对 DPanicLevel 进行堆栈跟踪
	disableCaller     bool                   // 不输出文件名和行数
	disableStacktrace bool                   // 不打印 trace 信息
	initialFields     map[string]interface{} // 加入一些初始的字段数据，比如项目名

	// EncoderConfig
	messageKey    string // 日志的 key name
	levelKey      string
	timeKey       string
	nameKey       string
	callerKey     string
	stacktraceKey string

	lineEnding     string // 换行符
	encodeTime     zapcore.TimeEncoder
	encodeLevel    zapcore.LevelEncoder
	encodeDuration zapcore.DurationEncoder
	encodeCaller   zapcore.CallerEncoder
	EncodeName     zapcore.NameEncoder

	// zap.Config
}

// ZapLogger 日志记录器
type ZapLogger struct {
	*zap.Logger
	sync.RWMutex
	opts      *options
	zapConfig zap.Config
	inited    bool
}

func newDefaultOptions() *options {
	return &options{

		isStdout: true,
		// zapConfig
		level:             zapcore.DebugLevel,
		development:       true,
		disableCaller:     false,
		disableStacktrace: false,
		initialFields:     nil,

		archive: archiveConfig{
			logFileDir:    "",
			errorFileName: "error.log",
			warnFileName:  "warn.log",
			infoFileName:  "info.log",
			debugFileName: "debug.log",
			maxSize:       11,
			maxBackups:    10,
			maxAge:        30,
		},

		// EncoderConfig
		messageKey:     "content",
		levelKey:       "level",
		timeKey:        "timestamp",
		nameKey:        "logger",
		callerKey:      "caller",
		stacktraceKey:  "stacktrace",
		lineEnding:     "\n",
		encodeTime:     zapcore.ISO8601TimeEncoder,
		encodeLevel:    zapcore.LowercaseLevelEncoder,
		encodeDuration: zapcore.StringDurationEncoder,
		encodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// NewLogger 初始化 logger
func NewLogger(mod ...ModOptions) {
	logger := &ZapLogger{}
	logger.Lock()
	defer logger.Unlock()
	if logger.inited {
		logger.Info("[NewLogger] logger Inited")
		return
	}

	logger.opts = newDefaultOptions()

	// 自定义覆盖默认配置
	for _, fn := range mod {
		fn(logger.opts)
	}

	logger.buildZapConfig()
	logger.buildEncoderConfig()

	logger.buildLogger()
	logger.inited = true

	if len(logger.zapConfig.InitialFields) > 0 {
		o := []zap.Option{}
		for k, v := range logger.zapConfig.InitialFields {
			o = append(o, zap.Fields(zap.Any(k, v)))
		}
		logger.Logger = logger.Logger.WithOptions(o...)
	}

	logger.Info("[NewLogger] success", zap.Bool("develop", logger.zapConfig.Development))

	SetGlobalLogger(logger)
}

// SetGlobalLogger 定义全局 logger
func SetGlobalLogger(l *ZapLogger) {
	globalLogger = l
}

// GetLogger returns the global logger
func GetLogger() *ZapLogger {
	return globalLogger
}

// 构建 zapConfig
func (l *ZapLogger) buildZapConfig() {

	if l.opts.development {
		l.zapConfig = zap.NewDevelopmentConfig()
		// l.zapConfig.EncoderConfig.EncodeTime = timeEncoder
	} else {
		l.zapConfig = zap.NewProductionConfig()
		// l.zapConfig.EncoderConfig.EncodeTime = timeUnixNano
	}

	l.zapConfig.Level.SetLevel(l.opts.level)

	if l.opts.disableCaller {
		l.zapConfig.DisableCaller = true
	}
	if l.opts.disableStacktrace {
		l.zapConfig.DisableStacktrace = true
	}
	if l.opts.archive.logFileDir == "" {
		l.opts.archive.logFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
		l.opts.archive.logFileDir += sp + "logs" + sp
	}
	if strings.HasSuffix(l.opts.archive.logFileDir, sp) {
		l.opts.archive.logFileDir = strings.TrimSuffix(l.opts.archive.logFileDir, sp)
	}
	if l.opts.initialFields != nil {
		l.zapConfig.InitialFields = l.opts.initialFields
	}
}

// 构建 EncoderConfig
func (l *ZapLogger) buildEncoderConfig() {
	if l.opts.messageKey != "" {
		l.zapConfig.EncoderConfig.MessageKey = l.opts.messageKey
	}
	if l.opts.levelKey != "" {
		l.zapConfig.EncoderConfig.LevelKey = l.opts.levelKey
	}
	if l.opts.timeKey != "" {
		l.zapConfig.EncoderConfig.TimeKey = l.opts.timeKey
	}
	if l.opts.nameKey != "" {
		l.zapConfig.EncoderConfig.NameKey = l.opts.nameKey
	}
	if l.opts.callerKey != "" {
		l.zapConfig.EncoderConfig.CallerKey = l.opts.callerKey
	}
	if l.opts.stacktraceKey != "" {
		l.zapConfig.EncoderConfig.StacktraceKey = l.opts.stacktraceKey
	}
	if l.opts.lineEnding != "" {
		l.zapConfig.EncoderConfig.LineEnding = l.opts.lineEnding
	}
	if l.opts.encodeTime != nil {
		l.zapConfig.EncoderConfig.EncodeTime = l.opts.encodeTime
	}
	if l.opts.encodeLevel != nil {
		l.zapConfig.EncoderConfig.EncodeLevel = l.opts.encodeLevel
	}
	if l.opts.encodeDuration != nil {
		l.zapConfig.EncoderConfig.EncodeDuration = l.opts.encodeDuration
	}
	if l.opts.encodeCaller != nil {
		l.zapConfig.EncoderConfig.EncodeCaller = l.opts.encodeCaller
	}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func timeUnixNano(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.UnixNano() / 1e6)
}
