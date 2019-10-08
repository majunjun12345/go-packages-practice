package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

/*
	https://www.jianshu.com/p/7414289c331f
	https://segmentfault.com/a/1190000010516906
*/

type Producer interface {
	MsgContent() string
}

type Receiver interface {
	Consumer([]byte) error
}

type RabbitMQ struct {
	Conn         *amqp.Connection
	Channel      *amqp.Channel
	Queue        string     // 队列名称
	RoutingKey   string     // key 名称
	ExchangeName string     // 交换机名称
	ExchangeType string     //交换机类型
	ProducerList []Producer // 生产者队列
	ReceiverList []Receiver // 消费者队列
}

func (r *RabbitMQ) Connect() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("connect to rabbitmq failed:", err)
		return
	}
	r.Conn = conn

	channel, err := conn.Channel()
	if err != nil {
		fmt.Println("create channel failed:", err)
	}
	r.Channel = channel
}

func (r *RabbitMQ) Close() {
	err := r.Channel.Close()
	if err != nil {
		fmt.Println("close channel failed:", err)
	}

	err = r.Conn.Close()
	if err != nil {
		fmt.Println("close connection failed:", err)
	}
}

func (r *RabbitMQ) Produce(producer Producer) error {
	fmt.Println("send message:", producer.MsgContent())
	fmt.Println(r.ExchangeName, r.RoutingKey)
	err := r.Channel.Publish(
		r.ExchangeName,
		r.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("mamengli"),
		},
	)
	return err
}

func (r *RabbitMQ) Receiver(receiver Receiver) error {
	/*
		prefetchCount: rabbitmq server 端收到确认消息后将会发送给 consumer 的消息数量
						如果 consumer 是 noAck，rabbitmq server 端将会忽视这个参数；
		prefetchSize: 基于字节数来控制
		global: 为 true 表示在这个连接上对现有和将来的所有消费者都有效
	*/
	// 在没有返回ack之前，最多只接收1个消息
	r.Channel.Qos(1, 0, true) // 获取消费通道,确保rabbitMQ一个一个发送消息,从而限制未ack的消息数量

	delivery, err := r.Channel.Consume(
		r.Queue, // queueName
		"",      // consumeName
		false,   // autoAck: 将autoAck设置为false，则需要在消费者每次消费完成消息的时候调用 Ack(false) 来告诉RabbitMQ该消息已经消费
		false,   // exclusive
		false,   // nolocal
		false,   // nowait
		nil,
	)
	if err != nil {
		fmt.Println("Consume err:", err.Error())
		return err
	}

	for msg := range delivery {
		// TODO: handler message
		// 直到数据处理成功后再返回，然后才会回复rabbitmq ack，否则需要重试
		fmt.Println("receive message:", string(msg.Body))
		for receiver.Consumer(msg.Body) != nil {
			fmt.Println("HandlerMsg 数据处理失败，将要重试")
			time.Sleep(1 * time.Second)
		}

		fmt.Println("HandlerMsg 数据处理成功")
		// 确认收到本条消息, multiple必须为false, 告诉RabbitMQ该消息可以删除
		msg.Ack(false)
	}

	return nil
}

// 消息处理函数可以注册进 RabbitMQ 对象中
func HandlerMsg(msg []byte) error {
	fmt.Println("handler msg:", string(msg))
	return nil
}

func (r *RabbitMQ) Init(routingKey, exchangeName, exchangeType, queueName string) error {

	// 注册交换机
	err := r.Channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,  // durable: 是否持久化
		false, // autoDelete: 是否自动删除
		false, // internal:是否为内部
		true,  // noWait: 是否非阻塞
		nil,
	)
	if err != nil {
		fmt.Println("注册交换机失败:", err)
		return err
	}

	// 注册队列
	_, err = r.Channel.QueueDeclare(
		queueName,
		true,  // durable: 是否持久化
		false, // autoDelete: 是否自动删除
		false, // exclusive: 排他性队列
		false, // noWait: 是否非阻塞
		nil,
	)
	if err != nil {
		fmt.Println("注册d队列失败:", err)
		return err
	}

	// 队列和交换机绑定
	err = r.Channel.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		true, // noWait: 是否非阻塞
		nil,
	)
	if err != nil {
		fmt.Println("队列和交换机绑定:", err)
		return err
	}

	r.RoutingKey = routingKey
	r.ExchangeName = exchangeName
	r.ExchangeType = exchangeType
	r.Queue = queueName
	return nil
}

// ----------------------------------------------------------------

type TestPro struct {
	Content string
}

// 实现发送者
func (t *TestPro) MsgContent() string {
	return t.Content
}

// 实现接收者
func (t *TestPro) Consumer(dataByte []byte) error {
	fmt.Println(string(dataByte))
	return nil
}

func main() {
	r := &RabbitMQ{}
	r.Connect()
	err := r.Init("test_route", "test_exchange", "direct", "test_queue")
	if err != nil {
		fmt.Println("init error", err)
	}
	testPro := &TestPro{
		Content: "mamengli",
	}

	err = r.Produce(testPro)
	if err != nil {
		fmt.Println("Produce message error", err)
	}
	r.Receiver(testPro)
}
