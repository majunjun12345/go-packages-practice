package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

const (
	mqhost = "amqp://guest:guest@10.231.20.48:5672/"
)

var (
	channel *amqp.Channel
)

func main() {
	InitChannel()
	ConsumeMessage()
}

// 创建连接和channel
func InitChannel() {
	var err error
	conn, err := amqp.Dial(mqhost)
	if err != nil {
		fmt.Println("dial err:", err.Error())
	}
	channel, err = conn.Channel()
	if err != nil {
		fmt.Println("create channel err:", err.Error())
	}
}

func ConsumeMessage() {
	msg, err := channel.Consume(
		"test_queue",
		"",    // 标签，自定义
		true,  // 收到消息后，不用手动发送确认消息给发布者，rabbitmq 客户端会自动帮我们做了
		false, // 是否指定一个消费者
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("consumer error:", err.Error())
	}
	for {
		select {
		case data := <-msg:
			fmt.Println(string(data.Body))
		}
	}
}
