package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// https://www.jianshu.com/p/ae72ad58ecb6
/*
	总结：
	超时语句如果放在 case 语句中，每次执行 select 后会重新计算时间。故一般将超时放置在 case 语句外；tick 也最好放置在 for select 语句外；
	放在外面可以和 default 共存，不然每次执行 default 语句都会重新计时；

	for select 语句中 接收 signal 信号 不能和 default 一起使用

	for select 中 break 只会跳出 select 语句，要想跳出 for，必须结合 goto 和 标签 或者 return；

	for select 语句中没有 switch case 的 fallthrough;
*/

func main() {
	/*
		退出主程序
	*/
	// ExitMain()

	/*
		优雅的退出守护进程
		守护进程即为后台执行的进程

		如果在守护进程中使用 for select 语句来监听信号，感觉监听不到，但是使用 for range 语句可以监听到，原因是使用了 default 语句，去掉 default 语句就行了；
		总结：default 语句 和 监听 signal 一起慎用！
	*/
	ExitGoroutine()
}

func ExitMain() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1) // 监听指定信号，默认所有
	fmt.Println("启动")
	s := <-c
	fmt.Printf("收到信号：%s, 退出程序", s)
}

func ExitGoroutine() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	a := make(chan int)

	go func() {
		/*
			可以监听到 signal 信号
		*/
		// for s := range c {
		// 	switch s {
		// 	case os.Interrupt:
		// 		fmt.Println("收到信号，退出程序")
		// 		ExitFunc()
		// 	}
		// }

		/*
			可以超时退出，也可以监听到 signal 退出
		*/
		// for {
		// 	select {
		// 	case <-c:
		// 		fmt.Println("s收到信号，退出程序")
		// 		ExitFunc()
		// 	case <-time.After(time.Second * 5):
		// 		fmt.Println("超时退出")
		// 		ExitFunc()
		// 	}
		// }

		/*
			当 signal 结合 default 使用时，监听不到信号
		*/

		timeout := time.After(time.Second * 5)
		tick := time.Tick(time.Second)

		for {
			select {
			case <-a:
				fmt.Println("收到结束信号")
				ExitFunc()
			case <-c:
				fmt.Println("收到 ctrl + c 退出程序")
				ExitFunc()
			case <-timeout:
				fmt.Println("timeout")
				ExitFunc()
			case <-tick:
				fmt.Println("tick")
			default:
				fmt.Println("working")
				time.Sleep(time.Second * 2)
			}
		}
	}()

	fmt.Println("启动程序")
	sum := 0
	for {
		sum++
		fmt.Println(sum)
		time.Sleep(time.Second * 10)
		a <- 1
	}
}

func ExitFunc() {
	fmt.Println("开始退出")
	fmt.Println("执行清理")
	fmt.Println("退出程序")
	os.Exit(0)
}
