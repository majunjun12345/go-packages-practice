package main

import (
	"fmt"
	"sync"
)

var unSafeInt int

type SafeInt struct {
	num int
	sync.Mutex
}

func main() {
	// TestLock()

	TestUnsafe()
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

func TestUnsafe() {
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			unSafeInt++
		}()
	}
	wg.Wait()
	fmt.Println(unSafeInt) // 915
}
