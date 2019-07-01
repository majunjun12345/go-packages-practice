package main

import (
	"log"
	"time"
)

/*
	channel 的使用场景：
		消息传递、消息过滤
		信号广播
		事件订阅与广播
		请求、响应转发
		任务分发
		结果汇总
		并发控制
		同步与异步
		...

	channel 的三种状态：
		nil：只进行了声明，未初始化的状态，或者手动赋值为 nil；
		active：正常读写状态；
		close：已关闭，不要认为关闭 channel 后，channel 的值是 nil；

	读 关闭的 chan，读到的是对应类型的 零值；
	将 select 中的某个 chan 置 nil，会阻塞，但是不会死锁；
*/

func main() {
	// testBufChan()

	testNilChan()
}

/*
	1. for range：
	当需要不断从 chan 读取数据时，使用 for range
	当 chan 关闭时，for 循环会自动退出，可以防止读取已经关闭的 chan，造成读到数据为通道所存储的数据类型的零值。
*/

/*
	2. v,ok := <-ch
	在 select 中
	ok的结果和含义：
		- true：读通道数据，不确定是否关闭（存在延时）。
			可能channel还有保存的数据（缓冲chan），但channel已关闭，这时候 ok 一直是 true。
			直到 chan 里面的数据全部被读取完，这时候 ok 才会是 false；
			也就是 chan 关闭后，依旧可以正常从 chan 中获取数据，知道数据全被被读出，这时候 ok 才是 false；
		- false：通道关闭，可以一直读取数据，但是都是类型零值。
*/
func testBufChan() {
	chan1 := make(chan int, 5)

	go func() {
		// defer close(chan1)
		for i := 0; i < 5; i++ {
			chan1 <- i
		}
		close(chan1)
		log.Println("close chan1")
	}()

	// for i := range chan1 {
	// 	time.Sleep(time.Second * 1)
	// 	log.Println("======:", i)
	// }

	for {
		num, ok := <-chan1
		log.Println("=====", num, ok)
		time.Sleep(time.Second)
	}
}

/*
	3. select 处理多个 channel
	只会处理最快获取数据的 channel
	select可以同时监控多个通道的情况，只处理未阻塞的case。当通道为nil时，对应的case永远为阻塞，无论读写。
	特殊关注：普通情况下，对nil的通道写操作是要panic的。
*/

func testNilChan() {
	chan1 := make(chan int)
	go func() {
		chan1 <- 1
	}()
	<-chan1
	// 两种情况都不行
	chan1 = nil
	<-chan1
	// chan1 <- 1

}

/*
	4. 使用单向 chan 控制读写权限
*/

/*
	5. 使用缓冲 chan 增强并发
*/
func testParrelChan() {
}

/*
	6. 加上超时操作
*/

/*
	7. 使用 close 关闭所有下游协程
*/

/*
	8. 使用 struct{} 作为信号
	这里为什么使用 空结构体 {} 呢?
	1. map 较长,浪费资源
	2. 表明这里不需要传值

	空结构体的声明赋值方式:
	声明: make(chan struct{})
	赋值: complete <- struct{}{}
*/
