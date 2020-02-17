package main

import (
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// https://www.cnblogs.com/saryli/p/11654665.html
// https://www.jianshu.com/p/05914c15b9a8

/*
  须知：
	broker: mqtt服务器，也称消息代理，介于消息发布者和消息订阅者之间

	token.Wait()：随处可见，为阻塞函数，操作完成后才返回

	SetDefaultPublishHandler：别看它带有 pub，其实是作用于 sud 侧的，且和 Subscribe 里的 handler 有优先级

	必须要保证订阅在前 发布在后，不然收不到消息
*/

/*
	- SetKeepAlive
		在上一课中，我们提到过 Broker 需要知道 Client 是否非正常地断开了和它的连接，以发送遗愿消息。
		实际上 Client 也需要能够很快地检测到它失去了和 Broker 的连接，以便重新连接。

		MQTT 协议是基于 TCP 的一个应用层协议，理论上 TCP 协议在丢失连接时会通知上层应用，
		但是 TCP 有一个半打开连接的问题（half-open connection）。这里我不打算深入分析 TCP 协议，需要记住的是，
		在这种状态下，一端的 TCP 连接已经失效，但是另外一端并不知情，它认为连接依然是打开的，
		它需要很长的时间才能感知到对端连接已经断开了，这种情况在使用移动或者卫星网络的时候尤为常见。

		仅仅依赖 TCP 层的连接状态监测是不够的，于是 MQTT 协议设计了一套 Keep Alive 机制。回忆一下，
		在建立连接的时候，我们可以传递一个 Keep Alive 参数，它的单位为秒，MQTT 协议中约定：在 1.5*Keep Alive 的时间间隔内，
		如果 Broker 没有收到来自 Client 的任何数据包，那么 Broker 认为它和 Client 之间的连接已经断开；
		同样地, 如果 Client 没有收到来自 Broker 的任何数据包，那么 Client 认为它和 Broker 之间的连接已经断开。
*/

// 创建全局mqtt publish消息处理 handler
var pushHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

// sub handler
var subscribeHandler = func(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Received message on topic: %s\n\tMessage: %s\n", message.Topic(), message.Payload())
}

func main() {

	// 设置连接参数和客户端id(这两个参数必须)，还可以 SetUsername("admin").SetPassword("public")
	opts := MQTT.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	opts.SetClientID("go-simple")
	// 设置连接超时
	opts.SetConnectTimeout(5 * time.Second)
	// 设置 handler
	opts.SetDefaultPublishHandler(pushHandler)
	// 心跳时间，单位秒
	opts.SetKeepAlive(30)
	// 创建客户端连接
	client := MQTT.NewClient(opts)

	// 客户端连接判断
	// token 用来指示操作是否完成，token.Wait()是个阻塞函数，只有在操作完成时才返回。token.WaitTimeout()会在操作完成后等待几毫秒后返回。
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 三个参数：订阅的 topic，订阅的qos质量，接受到匹配消息时的回调函数
	// 当回调函数为nil时，在接受到消息后会调用客户端的默认消息处理程序（如果设置），也就是说使用的先后顺序为 回调函数/SetDefaultPublishHandler
	if token := client.Subscribe("go-mqtt/sample", 0, subscribeHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// 四个参数：发布的 topic，消息的 qos 质量，是否保持消息连接的 bool，消息体 payload
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := client.Publish("go-mqtt/sample", 0, false, text)
		// token.Wait() 为阻塞函数，只有操作完成才返回
		token.Wait()
	}

	time.Sleep(1 * time.Second)

	// 取消订阅，可以取消订阅多个 topic
	if token := client.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// 断开连接，参数为等待的毫秒数，目的为等待已有工作的完成
	client.Disconnect(250)
}
