package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var unSafeInt int

type SafeInt struct {
	num int
	sync.Mutex
}

func main() {
	// TestLock()

	// TestUnsafe()

	// RWLock()

	TestCond1()

}

// ---------------------------- unsafe and lock

func TestUnsafe() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			unSafeInt++
		}()
	}
	wg.Wait()
	fmt.Println(unSafeInt) // 915
}

func TestLock() {
	count := SafeInt{}
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			count.Lock()
			defer wg.Done()
			defer count.Unlock()
			count.num += i
		}(i) // 这里要注意，有时候传递的是引用
	}
	wg.Wait()
	fmt.Println(count.num)
}

// -------------------------- RWLock
// 读写锁
var count int
var rw sync.RWMutex

func RWLock() {
	ch := make(chan struct{}, 10)

	for i := 0; i < 5; i++ {
		go TestRLock(i, ch)
	}
	for i := 0; i < 5; i++ {
		go TestWLock(i, ch)
	}

	for i := 0; i < 10; i++ {
		<-ch
	}
}

func TestRLock(n int, ch chan struct{}) {
	rw.RLock()
	defer rw.RUnlock()
	fmt.Printf("goroutine %v 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %v 读取结束，值为：%v\n", n, v)
	ch <- struct{}{}
}

func TestWLock(n int, ch chan struct{}) {
	rw.RLock()
	defer rw.RUnlock()
	fmt.Printf("goroutine %v 进入写操作...\n", n)
	count := rand.Intn(1000)
	fmt.Printf("goroutine %v 写入结束，新值为：%v\n", n, count)
	ch <- struct{}{}
}

// ------------------- cond 条件锁
// Cond是一个条件锁，就是当满足某些条件下才起作用的锁，有的地方也叫定期唤醒锁，有的地方叫条件变量conditional variable。
// 基于互斥锁，必须有互斥锁的支撑才能发挥作用

/*
	条件变量并不是用来保护临街区域和共享变量的，而是用来协调想要访问共享资源的一组线程；
	条件变量在这里最大的优势就是效率的提升，当共享资源不满足条件的时候，想要操作它的协程不需要轮询检查，等待通知即可；
*/

func TestCond1() {
	l := &sync.Mutex{}
	cond := sync.NewCond(l)

	for i := 0; i < 5; i++ {
		go func(i int) {
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(100)))
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait() // 阻塞当前协程，直到收到 cond 发来的通知
			fmt.Printf("this is the %v goroutine\n", i)
		}(i)
	}

	time.Sleep(2 * time.Second)
	cond.Signal()    // 下发一个通知给已经获取锁的 goroutine
	cond.Broadcast() // 下发通知给所有的 goroutine，排除 single 的 goroutine
	time.Sleep(2 * time.Second)
}
