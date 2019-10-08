package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const (
	TransExchangeName  = "uploadserver.trans"
	TransOSSRoutingKey = "oss"

	TransOSSQueueName = "uploadserver.trans.oss"
)

// 和 rabbitmq 建立连接
var conn *amqp.Connection

// 通过 conn 得到，主要用于消息的发布和接收
var Channel *amqp.Channel

func main1() {
	StartMQ()
	Send()
	Consume()
}

func StartMQ() {
	var err error

	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	Channel, err = conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
}

func Send() {
	if Channel == nil {
		fmt.Println("Channel is nil")
	}
	err := Channel.Publish(
		TransExchangeName,  // 交换机名
		TransOSSRoutingKey, // routing key
		false,              // false:找不到 queue，消息会被丢弃, true: 找不到 queue，直接返回消费者
		false,              // false: true：所有关联的 queue 上没有消费者，返回给消费者，有则立即投递
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("111"),
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("发送成功")
}

func Consume() {
	mesgChan, err := Channel.Consume(
		TransOSSQueueName,
		"menglima",
		true,  // 收到消息后，不用手动发送确认消息给发布者，rabbitmq 客户端会自动帮我们做了
		false, // 指定是否只有一个消费者
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("consume err:", err.Error())
		return
	}

	select {
	case mesg := <-mesgChan:
		fmt.Println("body:", string(mesg.Body))
	}

}
