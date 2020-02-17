package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/*
	分批次和异步写入：http://vearne.cc/archives/39275
*/

func main() {
	// CustomLog()

	conbineLumberjack()
}

// ---------------------------------------------------------------- default
func singleSimple() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()

	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "this is a url",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "this is a url")
}

// ---------------------------------------------------------------- custom
func CustomLog() {

	encodeConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "msg",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",                       // 只会对 warn 和 err 级别以上的日志进行堆栈跟踪
		LineEnding:     zapcore.DefaultLineEnding,          // 日志每行的分隔符
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,         // 输出的时间格式，ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,     // 执行消耗的时间转化成浮点型的秒
		EncodeCaller:   zapcore.ShortCallerEncoder,         // 以包/文件:行号 格式化调用堆栈, FullCallerEncoder
	}

	atom := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.Config{
		Level:            atom,                                                // 日志级别
		Development:      true,                                                // 是否是开发环境
		DisableCaller:    false,                                               // 是否禁止使用调用函数的文件名和行号来注释日志
		Encoding:         "json",                                              // 编码类型，json 和 console
		EncoderConfig:    encodeConfig,                                        // 生成格式的一些配置
		InitialFields:    map[string]interface{}{"serviceName": "spikeProxy"}, // 加入一些初始字段，比如项目名
		OutputPaths:      []string{"stdout", "./logs/spikeProxy.log"},         // 输出至日志文件的地址
		ErrorOutputPaths: []string{"stderr"},                                  // 将 error 记录到文件的地址
	}

	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("zap log 初始化失败: V%", err))
	}

	logger.Info("初始化成功")

	logger.Info("无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	logger.Warn("初始化失败")

	defer logger.Sync() // 输出所有 entries
}

// ----------------------------------------------------------------hook
func conbineLumberjack() {
	hook := lumberjack.Logger{
		Filename:   "./logs/spikeProxy.log",
		MaxSize:    100,   // 每个日志文件的大小 M
		MaxBackups: 3,     // 最大备份数
		MaxAge:     1,     // 文件最多保存天数
		Compress:   false, // 是否压缩
	}
	defer hook.Close()

	encodeConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "msg",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encodeConfig),                                            // 编码器配置                                   // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 将日志输出至标准输出和文件
		atomicLevel, // 日志等级
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	field := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	logger := zap.New(core, caller, development, field)

	for i := 0; i < 1; i++ {
		logger.Info("log success " + strconv.Itoa(i))
	}

	defer logger.Sync()
}
