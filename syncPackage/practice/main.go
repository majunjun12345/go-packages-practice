package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	curl()
}

var wg sync.WaitGroup

func curl() {
	body := make(chan string)

	urls := []string{
		"http://blog.csdn.net/wangshubo1989/article/details/77949252",
		"http://blog.csdn.net/wangshubo1989/article/details/77933654",
		"http://blog.csdn.net/wangshubo1989/article/details/77893561",
		// "http://www.baidu.com",
		// "http://www.baidu.com",
		// "http://www.baidu.com",
	}
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {

			defer wg.Done()

			response, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			content := response.Status
			if err != nil {
				panic(err)
			}

			body <- string(content)
		}(url)
	}

	go func() { // 这里有点 bug，有可能打印不全就退出了
		for text := range body {
			fmt.Println(text)
		}
	}()
	wg.Wait()
}
