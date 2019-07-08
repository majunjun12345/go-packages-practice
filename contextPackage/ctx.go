package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
	有点绕
*/
func main1() {
	ctx := context.Background()
	ctxWithCancel, cancelFunc := context.WithCancel(ctx)

	defer func() {
		fmt.Println("defer canceled by main")
		cancelFunc()
	}()

	go func() {
		sleepRandomTime("main", nil)
		cancelFunc()
		fmt.Println("Main Sleep complete. canceling context")
	}()

	doSomeWork(ctxWithCancel)
}

func sleepRandomTime(fromFunc string, ch chan int) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	fmt.Printf("sleep for:%v, %s\n", n, fromFunc)
	time.Sleep(time.Second * time.Duration(n))
	fmt.Printf("from func:%s\n", fromFunc)

	if ch != nil {
		ch <- n
	}
}

func sleepRandomTimeCtx(ctx context.Context, ch chan bool) {
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()

	sleepTimeChan := make(chan int)

	go sleepRandomTime("sleepRandomTimeCtx", sleepTimeChan)

	select {
	case <-ctx.Done():
		fmt.Println("sleepRandomContext: Time to return")
	case sleepTime := <-sleepTimeChan:
		fmt.Println("Slept for ", sleepTime, "ms")
	}
}

func doSomeWork(ctx context.Context) {
	ctxWithTimeOut, cancelFunc := context.WithTimeout(ctx, time.Duration(1500)*time.Millisecond)
	defer func() {
		fmt.Println("defer doSomeWork complete")
		cancelFunc()
	}()

	ch := make(chan bool)
	go sleepRandomTimeCtx(ctxWithTimeOut, ch)

	select {
	case <-ctx.Done():
		fmt.Println("doWorkContext: Time to return")
	case <-ch:
		fmt.Println("sleepRandomContext returned")
	}
}
