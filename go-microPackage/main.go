package main

import (
	"filestore-server/common"
	"fmt"
	"log"
	"time"

	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
)

func main() {
	go clock()
	startRPCservice()
}

func clock() {
	t := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-t.C:
			fmt.Println("ticker")
		}
	}
}

func startRPCservice() {
	service := micro.NewService(
		micro.Name("go.micro.service.transfer"), // 服务名称
		micro.RegisterTTL(time.Second*10),       // TTL指定从上一次心跳间隔起，超过这个时间服务会被服务发现移除
		micro.RegisterInterval(time.Second*5),   // 让服务在指定时间内重新注册，保持TTL获取的注册时间有效
		micro.Flags(common.CustomFlags...),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			mqhost := c.String("mqhost")
			if len(mqhost) > 0 {
				log.Println("custom mq address: " + mqhost)
				// config.UpdateMqHost(mqhost)
			}
		}),
	)

	service.Run()
}
