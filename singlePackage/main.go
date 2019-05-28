package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// https://www.jianshu.com/p/ae72ad58ecb6

func main() {
	/*
		退出主程序
	*/
	// c := make(chan os.Signal)
	// signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1) // 监听指定信号，默认所有
	// fmt.Println("启动")
	// s := <-c
	// fmt.Printf("收到信号：%s, 退出程序", s)

	/*
		优雅的退出守护进程
		守护进程即为后台执行的进程
	*/
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		for s := range c {
			switch s {
			case os.Interrupt:
				fmt.Println("收到信号，退出程序")
				ExitFunc()
			}
		}
	}()

	fmt.Println("启动程序")
	sum := 0
	for {
		sum++
		fmt.Println(sum)
		time.Sleep(time.Second)
	}
}

func ExitFunc() {
	fmt.Println("开始退出")
	fmt.Println("执行清理")
	fmt.Println("退出程序")
	os.Exit(0)
}
