package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	TimeFormat()
}

/*
	unix 和　unixNano
*/
func UnixTime() {
	t := time.Now()                            // 2019-06-11 22:17:44.282922603 +0800 CST m=+0.000105742
	fmt.Println(t.Unix())                      // 1560262709
	fmt.Println(t.UTC().Format(time.UnixDate)) // Tue Jun 11 14:19:11 UTC 2019

	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	fmt.Println(timestamp) //1560262815441619852  纳秒
	timestamp = timestamp[:10]
	fmt.Println(timestamp) // 1560262815　秒
}

/*
	时间格式化字符串转换
*/
func TimeFormat() {

	// ticker,每隔一段时间执行任务

	// timer　指定时间执行
	TimeChan()
}

func t() {
	now := time.Now()
	fmt.Println(now.Year())           // 2019
	fmt.Println(now.Month().String()) // June
	fmt.Println(now.Day())            // 12
	fmt.Println(now.Date())           // 2019 June 12
	fmt.Println(now.Hour())           // 8
	fmt.Println(now.Minute())         // 16
	fmt.Println(now.Second())         // 31

	// format
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 2019-06-12 08:20:12
	fmt.Println(time.Now().Format("2006-01-02#15:04:05")) // 2019-06-12#08:20:12
	fmt.Println(time.Now().Format("15:04:05"))            // 08:20:12
}

func TimeChan() {
	timeOut := time.After(time.Second * 10)
	ticker := time.Tick(time.Second * 2)
	// timer := time.NewTicker(time.Second * 10).C

	s := make(chan int)

	go func() {
		for {
			select {
			case <-s:
				fmt.Println("receive signal, stop!")
			case <-ticker:
				fmt.Println("tick")
			case <-timeOut:
				fmt.Println("time out")
			}
		}
	}()

	time.Sleep(time.Second * 15)
	s <- 3
	time.Sleep(time.Second * 2)
	fmt.Println("ending")
}
