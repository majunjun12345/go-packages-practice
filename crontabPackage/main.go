package main

import (
	"log"

	"github.com/robfig/cron"
)

func main() {
	i := 0
	c := cron.New()
	spec := "0*/1****"
	c.AddFunc(spec, func() {
		i++
		log.Println("excute per second", i)
	})
	c.Start()
	select {}
}
