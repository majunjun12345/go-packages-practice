package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

/*
	close(chan) 也往 chan 里面发送了信号,再加上一个 close 的功能;

	这里为什么使用 空结构体 {} 呢?
	1. map 较长,浪费资源
	2. 表明这里不需要传值

	空结构体的声明赋值方式:
	声明: chan struct{}
	赋值: complete <- struct{}{}
*/

var ShutDown bool

func main() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)

	complete := make(chan struct{})
	go LaunchProcessor(complete)

	// 这是一个阻塞机制, 直到接收到信号, 通知协程退出
	for {
		select {
		case <-sigChan:
			fmt.Println("shutdown")
			ShutDown = true
		case a, ok := <-complete: // 这里为什么能够收到信号?都没有向里面发送值
			fmt.Println("return:", a, ok)
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
			// complete <- struct{}{} // 这句可要可不要
			return
		}
	}

	fmt.Println("end work")
}
