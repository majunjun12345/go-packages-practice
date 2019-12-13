package main

import (
	"fmt"
	"sync"
)

/*
	https://www.jb51.net/article/150769.htm

	sync.Pool设计的目的是用来保存和复用临时对象，以减少内存分配，降低CG压力。
	pool 用于存储那些被分配了但是没有被使用，而未来可能会使用的值，以减小垃圾回收的压力，协程安全。
	原因：由于golang内建的GC机制会影响应用的性能，为了减少GC，golang提供了对象重用的机制，也就是sync.Pool对象池；

	get
	pool 会为每个 P 维护一个本地池，P 的本地池分为私有池 private 和共享池 shared
	私有池中的元素只能给本地 P 使用，访问时不用加锁。共享池中的元素会被其他 P 偷走，访问时需要加锁；
	get 会优先查找本地 private，再查找本地 shared，最后查找其他 P 的 shared，如果都没有，最后调用 New 函数获取新元素，如果 new 也没有，则返回 nil；

	put
	put 会优先将元素放入本地 private 池中；如果 private 不为空，则放本地 shared 池中；
	有趣的是，在放入池之前，该元素会有 1/4 被丢弃的可能；
*/

func main() {
	BasicalOption()
}

func BasicalOption() {
	// 初始化池, 构造一个对象就行，没必要一个具体的值
	pipe := sync.Pool{
		New: func() interface{} {
			return "hello"
		},
	}

	val := "hello world"
	// 向池中添加元素
	pipe.Put(val)

	// 获取池中已存的对象，否则创建一个
	first := pipe.Get().(string)
	fmt.Println(first)

	// 没有了，new 一个, 如果上述 new 未设置，则返回 nil
	second := pipe.Get().(string)
	fmt.Println(second)

	pipe.Put(1)
	fmt.Println(pipe.Get().(int))
}
