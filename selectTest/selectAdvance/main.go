package main

import (
	"fmt"
	"time"
)

/*
	nil 的通道永远阻塞
	如何跳出 for select
		在满足条件的 case 内,使用 return, 如果有结尾工作,尝试交给 defer
		在 select 外 for 内使用 break 跳出循环
		使用 goto
	select{} 永远阻塞
*/

func main() {
	testNilChan()
}

func testNilChan() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			chan1 <- i
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			chan2 <- i
		}
	}()

	go func() {
		res := nilChan(chan1, chan2)
		var total int
		for num := range res {
			total += num
		}
		fmt.Println(total)
	}()
	time.Sleep(100 * time.Second)
}

// nil 的通道永远阻塞
func nilChan(chan1, chan2 chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out) // 发送端close后, 接收端还能正常接收,通过第二个参数判断
		for {
			select {
			case x, open := <-chan1:
				if !open {
					chan1 = nil // 将 chan 置为 nil, 表示永久阻塞
					continue
				}
				out <- x
			case y, open := <-chan1:
				if !open {
					chan1 = nil
					continue
				}
				out <- y
			}

			if chan1 == nil && chan2 == nil {
				break
			}
		}

	}()

	return out
}
