package main

import (
	"fmt"
	"runtime"
)

/*
	runtime.Goexit()，调用此函数会立即使当前的goroutine的运行终止（终止协程），而其它的goroutine并不会受此影响。
	runtime.Goexit在终止当前goroutine前会先执行此goroutine的还未执行的defer语句。请注意千万别在主函数调用runtime.Goexit，因为会引发panic。
*/

func main() {
	// 1 : 0 1 2 3 4 a 5 6 7 8 9，让出 cpu 时间片输出 a
	// 2 随机，也有可能在输出 a 之前主进程就退出了，说明 主协程 和 go 携程 绑定了不同的 p
	runtime.GOMAXPROCS(1)
	exit := make(chan int)

	go func() {
		defer close(exit)
		go func() {
			fmt.Println("a")
		}()
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 4 {
			runtime.Gosched()
		}
	}
	<-exit
}
