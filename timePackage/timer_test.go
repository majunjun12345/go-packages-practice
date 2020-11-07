package main

import (
	"fmt"
	"log"
	"testing"
	"time"
)

/*
	timer: 单一事件

	如果在定时器前 stop() 其返回 true，过期后返回 false

	timer.Stop() 停止计时器, 并不会向 C 中发送数据，所以如果在定时器过期前执行 stop，定时任务不会执行;
*/
func TestTimerAndStop(t *testing.T) {
	var (
		timer *time.Timer
	)
	fmt.Println("==1==", time.Now().Format("2006-01-02 15:04:05"))
	timer = time.NewTimer(time.Second * 3)

	go func(timer *time.Timer) {
		for {
			select {
			case <-timer.C:
				fmt.Println("==2==", time.Now().Format("2006-01-02 15:04:05"))
			}
		}
	}(timer)
	time.Sleep(time.Second * 1)
	if timer.Stop() {
		fmt.Println("==3==", time.Now().Format("2006-01-02 15:04:05"))
	}

	time.Sleep(time.Second * 5)
}

/*
	已过期的定时器或者已停止的定时器，可以通过重置动作重新激活

	reset 只能作用于 已过期 或 stop 的 timer
	如果已经收到了 <- t.C 的值，表明已过期，那么 reset 可以直接使用

	最保险的方式是搭配 stop 使用，没过期就等到过期将 <- t.C 的值排干，过期了就直接使用
*/

func TestReset(t *testing.T) {
	var (
		timer *time.Timer
	)
	fmt.Println("==1==", time.Now().Format("2006-01-02 15:04:05"))
	timer = time.NewTimer(time.Second * 3)

	// timer.Stop() 为 true 表示 stop 成功，执行后续，如果为 false，表明已过期，需要将 <-timer.C 排干
	if !timer.Stop() {
		fmt.Println("==2==", time.Now().Format("2006-01-02 15:04:05"))
		<-timer.C
	}

	timer.Reset(time.Second * 3)
}

/*
	After()
	AfterFunc() 过期时执行回调
*/

// ----------------------------------------------------------------------------------------

// 使用场景，设定超时
func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second)

	select {
	case <-conn:
		timer.Stop()
		return true
	case <-timer.C: // 超时
		println("WaitChannel timeout!")
		return false
	}
}

// 延迟执行某个任务
func DelayFunction() {
	timer := time.NewTimer(5 * time.Second)

	select {
	case <-timer.C:
		log.Println("Delayed 5s, start to do something.")
	}
}

// ----------------------------------------------------------------------------------------

// select case 一个关闭的 chan，将会陷入无限循环中, 这种情况可以考虑用 for range 的方式
func TestChan(t *testing.T) {
	c := make(chan struct{})

	go func(c chan struct{}) {
		for {
			select {
			case <-c:
				fmt.Println("===========")
			}
		}
	}(c)
	close(c)
	time.Sleep(time.Second * 5)
}

// 如果只是 close 那么不会执行 for 里面的；如果只向 c 里面发送数据，那么 for 会一直阻塞；想要 for 正常退出，必须 发数据+close；
func TestRange(t *testing.T) {
	c := make(chan struct{})

	go func(c chan struct{}) {
		for d := range c {
			fmt.Println("======", d)
		}
		fmt.Println("===end===")
	}(c)
	c <- struct{}{}
	close(c)
	time.Sleep(time.Second * 5)
}
