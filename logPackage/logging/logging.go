package logging

import (
	"io"
	"os"

	logging "github.com/op/go-logging"
)

var ()

func InitLogger(logDir, name string, debug bool) (*logging.Logger, io.WriteCloser) {
	logger := logging.MustGetLogger(name)
	format := logging.MustStringFormatter(
		`%{color}[%{time:15:04:05.000}]%{level:.4s} %{id:03x}| %{shortfunc}: %{color:reset} %{message}`,
	)

	formatNoColor := logging.MustStringFormatter(
		`[%{time:15:04:05.000}]%{level:.4s} %{id:03x}| %{shortfunc}: %{message}`,
	)

	stdout := logging.NewLogBackend(os.Stdout, "", 0)

	file, err := NewRotateWriter(logDir, name)
	if err != nil {
		panic(err)
	}
	file.SetKeep(100)
	file.SetMax(1024 * 1024 * 20)

	logfile := logging.NewLogBackend(file, name, 0)

	stdoutFormatter := logging.NewBackendFormatter(stdout, format)
	logfileFormatter := logging.NewBackendFormatter(logfile, formatNoColor)

	// Only errors and more severe messages should be sent to logfile
	logfileLeveled := logging.AddModuleLevel(logfileFormatter)
	logfileLeveled.SetLevel(logging.INFO, "")
	if debug {
		logfileLeveled.SetLevel(logging.DEBUG, "")
	}

	// Set the backends to be used.
	logging.SetBackend(stdoutFormatter, logfileLeveled)

	return logger, file
}
