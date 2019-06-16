package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var cli *redis.Client

func main() {
	cli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	defer cli.Close()
	pong := cli.Ping().String() // 测试有没有连接上redis
	fmt.Println(pong)

	sub := cli.Subscribe("mychannel")
	msgChan := sub.Channel()
	for {
		select {
		case msg := <-msgChan:
			fmt.Println(msg.String())
			break
		}
	}
}
