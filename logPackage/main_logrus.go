package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Warn("A walrus appears")

	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "this is other field",
	})

	contextLogger.Info("hahah")
}
