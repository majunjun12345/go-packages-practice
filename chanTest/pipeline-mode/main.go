package main

import (
	"fmt"
	"sync"
)

func generateTask(taskChan chan int) {

	for i := 0; i < 8000; i++ {
		taskChan <- i
	}
	close(taskChan)
	fmt.Println("任务生产完成")
}

func generateSushu(taskChan chan int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range taskChan {
		// 判断是否是素数
		isSuShu := true
		for i := 2; i < t; i++ {
			if t%i == 0 {
				isSuShu = false
				break
			}
		}
		if isSuShu {
			resultChan <- t
		}
	}
}

func workPool(taskChan chan int, resultChan chan int, wg *sync.WaitGroup) {
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go generateSushu(taskChan, resultChan, wg)
	}
}

func main() {
	var (
		taskChan   = make(chan int, 8000)
		resultChan = make(chan int, 8000)
		// closeChan  = make(chan bool, 8)
		wg sync.WaitGroup
	)
	go generateTask(taskChan)
	workPool(taskChan, resultChan, &wg)

	wg.Wait()
	close(resultChan)

	for i := range resultChan {
		fmt.Println("=====", i)
	}
}
