package main

import (
	"fmt"
	"sync"
)

/*
	sync.Pool设计的目的是用来保存和复用临时对象，以减少内存分配，降低CG压力。
	pool 用于存储那些被分配了但是没有被使用，而未来可能会使用的值，以减小垃圾回收的压力，协程安全。
	原因：由于golang内建的GC机制会影响应用的性能，为了减少GC，golang提供了对象重用的机制，也就是sync.Pool对象池；
*/

func main() {
	BasicalOption()
}

func BasicalOption() {
	// 初始化池
	pipe := sync.Pool{
		New: func() interface{} {
			return "hello"
		},
	}

	val := "hello world"
	pipe.Put(val)

	first := pipe.Get().(string)
	fmt.Println(first)

	second := pipe.Get().(string)
	fmt.Println(second)

	pipe.Put(1)
	fmt.Println(pipe.Get().(int))
}
