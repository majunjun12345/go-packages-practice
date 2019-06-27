package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m      map[int]int
	wg     sync.WaitGroup
	rwLock sync.RWMutex
	lock   sync.Mutex
)

func main() {
	wLockTest()

	rLockTest()
}

func wLockTest() {
	m = make(map[int]int)
	wg.Add(100)

	for i := 0; i < 100; i++ {
		// go func() {
		// 	lock.Lock()
		// 	defer wg.Done()
		// 	defer lock.Unlock()
		// 	m[i] = i // 这里的 i 是引用，最终结果不如预期
		// }()

		go func(i int) {
			rwLock.Lock()
			defer wg.Done()
			defer rwLock.Unlock()
			m[i] = i // 这才是正确写法
		}(i)
	}
	wg.Wait()
	fmt.Println(m)
}

func rLockTest() {
	time.Sleep(2 * time.Second)
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) { // 这里的传值需要注意
			defer wg.Done()
			result := m[i]
			fmt.Println("result:", result)
		}(i)
	}
	wg.Wait()
}
