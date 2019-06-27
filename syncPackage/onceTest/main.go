package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	再次调用 once 方法将会被忽略掉，不会报错
	无论 once 里面的方法更换与否，都只会执行一次 once.do 函数

	一般用于给全局变量赋值，特别是读取配置文件的时候
*/
var once *sync.Once = &sync.Once{}

func main() {

	fmt.Println("begin test")

	go test1(once)
	go test1(once)
	time.Sleep(time.Second * 1)
	go test2()
	time.Sleep(time.Second * 3)
}

// do 里面函数相同
func test1(one *sync.Once) {
	fmt.Println("======begin")
	one.Do(func() { // 再次调用将会忽略
		fmt.Println("======once 1")
	})
	fmt.Println("======end")
}

// do 里面函数不同
func test2() {
	fmt.Println("begin2")
	once.Do(op)
	fmt.Println("++++++end2")
}

func op() {
	fmt.Println("once 2")
}
