package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// server()

	testTimeout()
}

func server() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// monitor
		go func() {
			// 这种写法有点意思
			for range time.Tick(time.Duration(1) * time.Second) {
				select {
				case <-r.Context().Done():
					fmt.Println("req is outgoing")
					return
				default:
					fmt.Println("req is processing", time.Now().Unix())
				}
			}
		}()
		time.Sleep(time.Millisecond * 3500)
		w.Write([]byte("hello"))
	})
	http.ListenAndServe(":8081", nil)
}

func testTimeout() {
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Duration(3)*time.Second)

	// 	虽然超时后会自动关闭，但 cancel 代码仍建议加上
	defer cancel()

	select {
	case <-time.After(time.Duration(2) * time.Second):
		fmt.Println("overslept")
	case <-ctxWithTimeout.Done():
		fmt.Println(ctx.Err())
	}
}
