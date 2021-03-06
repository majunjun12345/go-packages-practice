package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

/*
	- NewPool
		NewPool 初始化
		Submit 利用 for 循环提交任务及执行器
*/

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	defer ants.Release()

	// 任务次数
	runTimes := 100

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}

	pool, err := ants.NewPool(10)
	if err != nil {
		panic(err)
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = pool.Submit(syncCalculateSum)
	}
	wg.Wait()

	fmt.Printf("===running goroutines: %d\n", ants.Running())
	fmt.Printf("===finish all tasks.\n")

	// ----------------------------------------------------------------

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	// p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
	// 	myFunc(i)
	// 	wg.Done()
	// })
	// defer p.Release()
	// // Submit tasks one by one.
	// for i := 0; i < runTimes; i++ {
	// 	wg.Add(1)
	// 	// int32(i) 是作为上面 NewPoolWithFunc 函数的参数
	// 	_ = p.Invoke(int32(i))
	// }
	// wg.Wait()
	// fmt.Printf("running goroutines: %d\n", p.Running())
	// fmt.Printf("finish all tasks, result is %d\n", sum)
}
