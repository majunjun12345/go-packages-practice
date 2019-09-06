package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	f, _ := os.OpenFile("logs/log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0766)

	out := io.MultiWriter(f, os.Stdout)
	log.SetOutput(out)
	log.SetLevel(logrus.DebugLevel)
	Logger = log
}

func main() {
	Logger.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	Logger.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Warn("A walrus appears")

	contextLogger := Logger.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other":  "this is other field",
	})

	contextLogger.Info("hahah")
}
