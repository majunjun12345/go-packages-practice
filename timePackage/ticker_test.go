package main

import (
	"fmt"
	"testing"
	"time"
)

/*
	ticker 和 sleep 的区别：
	tick的实现 使用了一个协程来进行定时 任务执行的时间会对其实际间隔时间产生影响，它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者
	tick的(上一个begin到下一个begin时间) = max (定时间隔时间, 任务执行时间)
*/

func TestTicker(t *testing.T) {
	var (
		ticker *time.Ticker
		n      int
	)

	ticker = time.NewTicker(time.Second * 3)

	for {
		select {
		case <-ticker.C:
			fmt.Println("begin", time.Now().Format("2006-01-02_15:04:05"))
			time.Sleep(time.Second * 1)
			fmt.Println("end", time.Now().Format("2006-01-02_15:04:05"))

			n++

			// break 只会跳出 select，不会跳出 for
			if n == 3 {
				ticker.Stop()
				goto END
			}
		}
	}
END:
}
