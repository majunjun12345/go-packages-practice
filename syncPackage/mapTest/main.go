package main

import (
	"fmt"
	"sync"
)

/*
	golang 中的 map 不是线程安全的

	sync.Map 线程安全，无需初始化即可使用
	sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用。Store 表示存储，Load 表示获取，Delete 表示删除
	使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值。Range 参数中的回调函数的返回值功能是：需要继续迭代遍历时，返回 true；终止迭代遍历时，返回 false
*/

func main() {
	// fatal error: concurrent map writes
	// TestOriginMap()

	TestSyncMap()
}

func TestSyncMap() {
	var syncMap sync.Map

	// 存储值
	syncMap.Store("name", "mamengli")
	syncMap.Store("age", 21)
	syncMap.Store("married", false)

	// 获取值
	name, ok := syncMap.Load("name")
	if ok {
		fmt.Println(name)
	}

	// 删除
	syncMap.Delete("name")

	// 遍历
	syncMap.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
	// 如果 map 中存在 age 给定的值，那么返回该值；如果不存在, 存储该键值对并返回值；如果 map 中存在 key，返回 true，不存在则返回 false；
	fmt.Println(syncMap.LoadOrStore("age", "22"))
	fmt.Println(syncMap.LoadOrStore("married", false))
	fmt.Println(syncMap.LoadOrStore("home", "xiantao"))

	for i := 0; i < 100; i++ {
		go func() {
			syncMap.Store(i, i)
		}()
	}
	fmt.Println(syncMap)
}

var m map[int]int = make(map[int]int)

func TestOriginMap() {

	for i := 0; i < 10; i++ {
		go func() {
			m[i] = i
		}()
	}
	fmt.Println(m)
}
