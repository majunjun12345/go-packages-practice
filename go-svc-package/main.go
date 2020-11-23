package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/judwhite/go-svc/svc"
)

func main() {
	prg := &Program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}

type Program struct {
}

// Init 程序运行前的初始化
func (p *Program) Init(e svc.Environment) error {
	fmt.Println("Initializing server...")
	fmt.Println(e.IsWindowsService())
	return nil
}

// Start 开启进程，不能阻塞s
func (p *Program) Start() error {
	fmt.Printf("程序已经start\n")
	go func() {
		tik := time.NewTicker(3 * time.Second)
		for {
			c := <-tik.C
			fmt.Printf("当前时间为:%s \n", c.Format("2006-01-02 15:04:05"))
		}

	}()
	return nil
}

// Stop 释放资源
func (p *Program) Stop() error {
	log.Printf("程序已经stop\n")
	return nil
}
