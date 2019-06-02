package main

import (
	"io"
	"uploader2/logging"

	gologging "github.com/op/go-logging"
)

/*
	backaend 就是输出至不同位置,终端 或 文件

	流程:
	1. GetLogger
	MustGetLogger 和 GetLogger 功能一样,当不能写日志时会 panic
	2. format
	可以定义多个 format,不同的 backend 会形成不同的 format
	3. backend
	可以定义 stdout 和 file,将日志在文件和终端中显示

*/

var (
	logger  *gologging.Logger
	logFile io.WriteCloser
)

func init() {
	logger, logFile = logging.InitLogger("./logs", "uploader", true)
}

func main() {
	logger.Info("mamengli")
	logger.Warning("menglima")

}
