package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

var (
	limiter *rate.Limiter
)

func init() {
	/*
		初始化 limiter，r: 每秒可以向 token 桶中产生多少令牌, 还可以通过 rate.Every 指定放置 token 的间隔； b: token 桶中的初始容量

	*/
	limiter = rate.NewLimiter(5, 10)
}

func main() {
	// Allow()
	Wait()
}

func Allow() {

	for {
		if limiter.AllowN(time.Now(), 1) {
			fmt.Println("=====consume=====")
			continue
		}
		fmt.Println("=====leak token=====")
		time.Sleep(time.Millisecond * 100)
	}
}

func Wait() {
	go func() {
		for {
			// TODO: 错误的写法，defer 不能用在死循环里面，这里只是演示
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()
			if err := limiter.WaitN(ctx, 1); err != nil {
				fmt.Println("xxx", err.Error())
				time.Sleep(time.Millisecond * 100)
				continue
			}
			fmt.Println("=====consume=====")
		}
	}()
	select {}
}
