package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

/*
	https://www.jianshu.com/p/7414289c331f
	https://segmentfault.com/a/1190000010516906

	- 持久性
		消息的持久性(发送时)
		队列的持久性
		exchange 的持久化

	- 断线重连
		客户端需要负责断线重连的逻辑是很重要的，因为有可能客户端和HAProxy的连接是正常的，
		但是HAProxy和rabbitmq的链接因为网络波动断开了，那么这个时候客户端其实是没有工作的，
		并且会在rabbitmq中不断积累消息。
	- Ack
		- 消息发送确认
			发送确认分为两步，一是确认是否到达交换器，二是确认是否到达队列

		- 消费接收确认
			消费消息的时候可以指定 autoAck 为 true 或 false
			- 当为 true 时
				RabbitMQ会自动把发送出去的消息标记为确认，
				然后从内存或者磁盘中移除，而不管消费者有没有收到消息，或者消息有没有处理成功。
			- 当为 false 时
				RabbitMQ会等待显示的恢复确认信号之后才从内存或者磁盘中移除消息（实质上是先打上删除标记，之后再删除）
				如果当前 consumer 挂了，未经确认的消息还会被投送给其他 consumer;
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
	QueueName    string     // 队列名称
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
	// channel.Confirm(true)
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
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte("mamengli"),
		},
	)
	return err
}

func (r *RabbitMQ) Receiver(receiver Receiver) error {
	defer r.Close()
	/*
		prefetchCount: rabbitmq server 端收到确认消息后将会发送给 consumer 的消息数量
						如果 consumer 是 noAck，rabbitmq server 端将会忽视这个参数；
		prefetchSize: 基于字节数来控制
		global: 为 true 表示在这个连接上对现有和将来的所有消费者都有效
	*/
	// 在没有返回ack之前，最多只接收1个消息
	r.Channel.Qos(1, 0, true) // 获取消费通道,确保rabbitMQ一个一个发送消息,从而限制未ack的消息数量

	delivery, err := r.Channel.Consume(
		r.QueueName, // queueName
		"",          // consumeName
		false,       // autoAck: 将autoAck设置为false，则需要在消费者每次消费完成消息的时候调用 Ack(false) 来告诉RabbitMQ该消息已经消费
		false,       // exclusive
		false,       // nolocal
		false,       // nowait
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
		if receiver.Consumer(msg.Body) != nil {
			fmt.Println("HandlerMsg 数据处理失败，将要重试")
			// msg.Ack(true)
			time.Sleep(1 * time.Second)
		} else {
			fmt.Println("HandlerMsg 数据处理成功")
			// 确认收到本条消息, multiple必须为false, 告诉RabbitMQ该消息可以删除
			msg.Ack(false)
		}
	}
	return nil
}

// 消息处理函数可以注册进 RabbitMQ 对象中
func HandlerMsg(msg []byte) error {
	fmt.Println("handler msg:", string(msg))
	return nil
}

// 初始化队列交换机
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
		fmt.Println("注册队列失败:", err)
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
	r.QueueName = queueName
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
		Content: "mamengli1",
	}
	// testPro2 := &TestPro{
	// 	Content: "mamengli2",
	// }
	// testPro3 := &TestPro{
	// 	Content: "mamengli3",
	// }

	err = r.Produce(testPro)
	// err = r.Produce(testPro2)
	// err = r.Produce(testPro3)
	if err != nil {
		fmt.Println("Produce message error", err)
	}
	r.Receiver(testPro)
}
