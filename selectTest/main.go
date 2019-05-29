package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

var ShutDown bool

func main() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)

	complete := make(chan struct{})
	go LaunchProcessor(complete)
	for {
		select {
		case <-sigChan:
			fmt.Println("shutdown")
			ShutDown = true
		case <-complete:
			fmt.Println("return")
			return
		}
	}
}

func LaunchProcessor(complete chan struct{}) {
	defer func() {
		close(complete)
	}()

	fmt.Println("start work")

	for i := 0; i < 5; i++ {
		fmt.Println("doing work")
		time.Sleep(time.Second * 2)

		if ShutDown {
			fmt.Println("kill work")
			return
		}
	}

	fmt.Println("end work")
}
