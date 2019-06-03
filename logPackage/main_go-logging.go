package main

import (
	"os"

	goLogging "github.com/op/go-logging"
)

/*
	backaend 就是输出至不同位置, 终端 或 文件

	流程:
	1. GetLogger  奇怪的是下面几个都不需要使用到，最后直接 log 日志就行
		MustGetLogger 和 GetLogger 功能一样,当不能写日志时会 panic
	2. NewLogBackend    原始
		定义输出位置，可以定义单个，也可以定义多个，一般是 file 和 stdout
	3. MustStringFormatter     原始
		可以定义多个 format,不同的 backend 会形成不同的 format
		time level pid shortfile shortfunc message
	4. NewBackendFormatter 原始
		将 backend 和 format 结合
	5. AddModuleLevel + SetLevel 原始
		设置输出至不同 backend 的日志等级
	6. SetBackend
		Set the backends to be used
*/

// var (
// 	logger  *gologging.Logger
// 	logFile io.WriteCloser
// )

// func init() {
// 	logger, logFile = logging.InitLogger("./logs", "uploader", true)
// }

func main1z() {
	// logger.Info("mamengli")
	// logger.Warning("menglima")

	// 不能直接 logger := goLogging.Logger

	file, _ := os.OpenFile("testGoLogging.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	logger := goLogging.MustGetLogger("testLog")

	logFile := goLogging.NewLogBackend(file, "fileOut", 0)
	logStdout := goLogging.NewLogBackend(os.Stdout, "stdOut", 0)

	fileFormat := goLogging.MustStringFormatter(
		"%{time:15:04:05.000} %{level:.4s} %{pid} %{id:03x} %{shortfile}] %{shortfunc}: %{message}",
	)
	stdoutFormat := goLogging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{level:.4s} %{id:03x} %{color:reset}%{shortfunc}: %{message}`,
	)

	fileBackend := goLogging.NewBackendFormatter(logFile, fileFormat)
	stdBackend := goLogging.NewBackendFormatter(logStdout, stdoutFormat)

	fileLevelBackend := goLogging.AddModuleLevel(fileBackend)
	fileLevelBackend.SetLevel(goLogging.DEBUG, "")
	stdoutLevelBackend := goLogging.AddModuleLevel(stdBackend)
	stdoutLevelBackend.SetLevel(goLogging.ERROR, "")

	goLogging.SetBackend(fileLevelBackend, stdoutLevelBackend) // 封装 MultiLogger

	logger.Info("menglima")
}
