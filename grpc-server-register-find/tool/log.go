package tool

import "testGoScripts/zaplog"

// InitLog 初始化日志
func InitLog() {
	var (
		l zaplog.Level
	)
	zaplog.NewLogger(
		zaplog.SetInitialFields("serviceName", "xxxhello-worldxxx"),
		zaplog.SetDevelopment(true),           // 环境: dev pro
		zaplog.SetLevel(l.Unmarshal("debug")), // log level
		zaplog.SetLogFileDir("/tmp"),          // log file directory
		zaplog.SetMaxSize(10),                 // 每个文件切分大小
		zaplog.SetMaxAge(7),                   // 文件最大生命周期
		zaplog.SetMaxBackups(10),              // 备份数
		zaplog.SetNoStacktrace(false),         // 不打印出 trace 信息
	)
}
