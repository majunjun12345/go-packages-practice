package main

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	/*
		初始化 limiter，r: 每秒可以向 token 桶中产生多少令牌, 还可以通过 rate.Every 指定放置 token 的间隔； b: token 桶中的初始容量

	*/
	limiter := rate.NewLimiter(rate.Every(100*time.Millisecond), 10)

	for {
		if limiter.AllowN(time.Now(), 1) {
			fmt.Println("=====consume=====")
			continue
		}
		fmt.Println("=====leak token=====")
		time.Sleep(time.Millisecond * 100)
	}
}
