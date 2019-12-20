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

	SetDefaultPublishHandler：别看它带有 pub，其实是在 sub 侧设置的，且和 Subscribe 里的 handler 有优先级

	必须要保证订阅在前 发布在后，不然收不到消息
*/

// 创建全局mqtt publish消息处理 handler
var pushHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

// sub handler
var subscribeHandler = func(client MQTT.Client, message MQTT.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

func main() {

	// 设置连接参数和客户端id(这两个参数必须)，还可以 SetUsername("admin").SetPassword("public")
	opts := MQTT.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	opts.SetClientID("go-simple")
	// 设置连接超时
	opts.SetConnectTimeout(60 * time.Second)
	// 设置 handler
	opts.SetDefaultPublishHandler(pushHandler)
	// 创建客户端连接
	c := MQTT.NewClient(opts)

	// 客户端连接判断
	// token 用来指示操作是否完成，token.Wait()是个阻塞函数，只有在操作完成时才返回。token.WaitTimeout()会在操作完成后等待几毫秒后返回。
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 三个参数：订阅的 topic，订阅的qos质量，接受到匹配消息时的回调函数
	// 当回调函数为nil时，在接受到消息后会调用客户端的默认消息处理程序（如果设置），也就是说使用的先后顺序为 回调函数/SetDefaultPublishHandler
	if token := c.Subscribe("go-mqtt/sample", 0, subscribeHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// 四个参数：发布的 topic，消息的 qos 质量，是否保持消息连接的 bool，消息体 payload
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish("go-mqtt/sample", 0, false, text)
		// token.Wait() 为阻塞函数，只有操作完成才返回
		token.Wait()
	}

	time.Sleep(1 * time.Second)

	// 取消订阅，可以取消订阅多个 topic
	if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// 断开连接，参数为等待的毫秒数，目的为等待已有工作的完成
	c.Disconnect(250)
}
