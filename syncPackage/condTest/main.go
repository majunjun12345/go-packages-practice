package main

import (
	"log"
	"sync"
	"time"
)

/*
	并发的关键在于同步和通信,go 中解决此类问题的方法有 sync channel atomic 等
	互斥锁是同步

	cond 是条件锁:当满足某些条件下才起作用的锁,基于互斥锁
	条件变量并不是用来保护临界区和共享资源的,主要用来协调想要访问共享资源的那些线程,当共享资源的状态发生变化时, 它可以用来通知被互斥锁阻塞的线程;

	条件变量最大的优势在于效率方面的提升, 当共享资源的状态不满足条件时, 想操作它的线程不用循环往复的做检查, 等待通知就行了;
*/

var cond *sync.Cond

func init() {
	cond = sync.NewCond(&sync.Mutex{})
}

func test(i int) {
	cond.L.Lock()         // 获取锁
	defer cond.L.Unlock() // 释放锁

	cond.Wait() // 等待通知, 暂时阻塞
	log.Println(i)
	time.Sleep(time.Second * 1)
}

func CondTest1() {

	for i := 0; i < 40; i++ { // 这里 cond 同时开了 40 个锁, 按顺序等待通知
		go test(i)
	}

	log.Println("start all")
	time.Sleep(time.Second * 3)
	log.Print("broadcast")
	cond.Signal() // 下发一个通知给已经获取锁的 goroutine, 最靠前的那一个 goroutine 才能拿到这个通知
	time.Sleep(time.Second * 3)
	cond.Signal() // 再下发一个通知给已经获取锁的 goroutine
	time.Sleep(time.Second * 3)
	cond.Broadcast() // 广播给所有等待的 goroutine
	time.Sleep(time.Second * 60)
	log.Println("Done!")
}
func main() {
	CondTest1()
}
