package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron"
)

/*
	`*` 秒 0-59
	`*` 分 0-59
	`*` 时 0-23
	`*` 日 1-31
	`*` 月 1-12
	`*` 星期 0-6 非必须

	/：表示增长间隔，默认是 1，3-59/15 表示第3秒开始执行一次，之后每隔15s执行一次，也可以表示为：3/15
	,：用于枚举，3,8,26 表示 第 3 8 26 秒各执行一次
	-：表示范围，3-25，表示 3 到 25 秒每秒执行一次，包括 3 和 25
*/

func main() {

	// RunSingleJob()
	RunJobs()
}

// 单个任务
func RunSingleJob() {
	i := 0
	c := cron.New()
	spec := "*/5 * * * *" // 每 5s 执行一次
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()
	select {}
}

// 多个任务
type Job1 struct {
}

func (j Job1) Run() {
	fmt.Println("test job1")
}

type Job2 struct {
}

func (j Job2) Run() {
	fmt.Println("test job2")
}

func RunJobs() {
	i := 0
	c := cron.New()

	spec1 := "*/1 * * * *"
	spec2 := "*/5 * * * *"
	spec3 := "*/10 * * * *"

	c.AddFunc(spec1, func() {
		i++
		fmt.Println("test jobfunc", i)
	})

	c.AddJob(spec2, Job1{})
	c.AddJob(spec3, Job2{})

	c.Start() // 启动任务

	select {}
}
