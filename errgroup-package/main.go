package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
	1. errgroup.Group在出现错误或者等待结束后都会调用 Context对象 的 cancel 方法同步取消信号。
	2. 只有第一个出现的错误才会被返回，剩余的错误都会被直接抛弃。
*/

func main() {
	demo2()
}

func demo1() {
	var eg errgroup.Group
	for i := 0; i < 100; i++ {
		j := i

		// 并发执行
		eg.Go(func() error {
			time.Sleep(2 * time.Second)
			if j > 90 {
				fmt.Println("Error:", j)
				return fmt.Errorf("Error occurred: %d", j)
			}
			fmt.Println("End:", j)
			return nil
		})
	}

	// 遇到错误后会继续执行其他协程，但是只有第一个 err 被返回
	if err := eg.Wait(); err != nil {
		fmt.Println("===============", err)
	}
}

func demo2() {
	eg, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < 100; i++ {
		j := i
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				fmt.Println("Canceled:", i)
				return nil
			default:
				if j > 90 {
					fmt.Println("Error:", j)
					return fmt.Errorf("Error occurred: %d", j)
				}
				fmt.Println("End:", j)
				return nil
			}
		})
	}

	// 遇到错误后，会取消执行其他 goroutine
	if err := eg.Wait(); err != nil {
		fmt.Println("===============", err)
	}
}
