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

	chan 被 close 后，相当于又往 chan 里面发送了其类型的零值，所以接收端在 close 后还会收到值，第二个参数变成 false；
	如果 chan 在 close 的时候没有被接收，将会导致死锁；

	用 for range 遍历 chan，在其 close 后会自动退出；
	普通方式通过第二个值判断是否 close；
*/

func main() {
	// testNilChan()

	testClose()
}

func testNilChan() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		defer func() {
			close(chan1)
			close(chan2)
		}()
		chan1 <- 1
		chan2 <- 2
	}()

	go func() {
		res := nilChan(chan1, chan2)
		var total int
		// 以下两种获取方式都可以
		// for num := range res {
		// 	total += num
		// }

		for {
			num, ok := <-res
			fmt.Println("11111:", num, ok)
			if !ok {
				break
			}
			total += num
		}
		fmt.Println(total)
	}()

	time.Sleep(10 * time.Second)
}

func nilChan(chan1, chan2 chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out) // 发送端close后, 接收端还能正常接收, 通过第二个参数判断
		for {
			select {
			case x, open := <-chan1: // chan1 被 close 后， x 对应 chan 类型的零值，open 为 false（也就是 close 后还可以接收一次）
				fmt.Println("aaaaa", x, open)
				if !open {
					chan1 = nil // 将 chan 置为 nil, 将会永久阻塞
				}
				out <- x
			case y, open := <-chan2:
				fmt.Println("bbbbb", y, open)
				if !open {
					chan2 = nil
				}
				out <- y
			}

			if chan1 == nil && chan2 == nil {
				fmt.Println("nil")
				break
			}
		}
	}()
	return out
}

func testClose() {
	ch := make(chan int)

	ch <- 1
	close(ch) // 死锁
}
