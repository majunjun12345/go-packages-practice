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
	PreWork()
	PublishMessage([]byte("helloworld"))
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

func PreWork() {
	// 创建 exchange
	err := channel.ExchangeDeclare(
		"test_exchange", // exchange name
		"direct",        // 工作模式
		true,            // 持久化
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("create exchange fail:", err.Error())
	}

	// 创建 queue
	channel.QueueDeclare(
		"test_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("create queue fail:", err.Error())
	}

	// exchange 绑定 queue
	err = channel.QueueBind(
		"test_queue",
		"test_bind",
		"test_exchange",
		false,
		nil,
	)
	if err != nil {
		fmt.Println("exchange bind queue fail:", err.Error())
	}
}

func PublishMessage(msg []byte) {
	channel.Publish(
		"test_exchange",
		"test_bind",
		false,
		false,
		amqp.Publishing{
			Headers:     amqp.Table{},
			ContentType: "text/plain",
			Body:        msg,
			// Expiration:  "6000", 消息过期时间
		},
	)
}
