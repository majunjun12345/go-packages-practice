package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
	channel 用完后必须关闭，不然会导致内存泄漏

	这个例子的关注点有四个：
		1. 一个无缓冲 chan 用来不断接收数据 -- 生产
		2. 开 n 个协程从 chan 里面读取数据  -- 消费
		3. 将消费的结果塞到无缓冲 chan
		4. 阻塞，从无缓冲 chan 获取全部协程结果汇总

	重要：
		何时关闭 chan，因为 chan 不关闭，range 不会退出, 造成死锁，这里也可以堪称是一种阻塞机制
*/

func Add() {
	begin := time.Now()

	wg := sync.WaitGroup{}
	num := make(chan int)
	sum := make(chan int)

	go func() {
		for i := 0; i < 10000000; i++ {
			num <- i
		}
		defer close(num)
	}()

	n := runtime.NumCPU()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			total := 0
			for i := range num {
				total += i
			}
			fmt.Println(total)
			sum <- total
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(sum)
	}()

	t := 0
	for n := range sum {
		t += n
	}

	cost := time.Since(begin)
	fmt.Println("cost:", cost)
	fmt.Println(t)
}

// 奇怪，这个快将近 100 倍
func add() {
	begin := time.Now()
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	cost := time.Since(begin)
	fmt.Println("cost1:", cost)
	fmt.Println(sum)
}
