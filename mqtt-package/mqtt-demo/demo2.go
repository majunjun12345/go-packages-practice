package main

import (
	"flag"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// 创建全局mqtt publish消息处理 handler
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("【Pub】 Client Topic : %s \n", msg.Topic())
	// fmt.Printf("Pub Client msg : %s \n", msg.Payload())
}

// 创建全局mqtt sub消息处理 handler
var messageSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("【Sub】 Client Topic : %s \n", msg.Topic())
	// fmt.Printf("Sub Client msg : %s \n", msg.Payload())
}

// 连接失败数
var failNums = 0

/*
 创建客户端连接
*/
func main() {
	flag.Parse()

	go mqttConnSubMsgTask(0)

	go mqttConnPubMsgTask(0)

	select {}
}

/*
 连接任务和发布消息方法
*/
func mqttConnPubMsgTask(taskId int) {
	time.Sleep(time.Second)
	// 设置连接参数
	clinetOptions := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	// 设置客户端ID
	clinetOptions.SetClientID(fmt.Sprintf("go Publish client example： %d-%d", taskId, time.Now().Unix()))

	clinetOptions.SetAutoReconnect(true)
	// 设置连接超时
	clinetOptions.SetConnectTimeout(time.Duration(60) * time.Second)
	// 创建客户端连接
	client := mqtt.NewClient(clinetOptions)

	// 客户端连接判断
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		failNums++
		fmt.Printf("[Pub] mqtt connect error, taskId: %d, fail_nums: %d, error: %s \n", taskId, failNums, token.Error())
		return
	}

	// for i := 0; i < 3; i++ {
	text := fmt.Sprintf("this is test msg #%d ! from task :%d", 1, taskId)
	// 发布消息
	token := client.Publish("go-test-topic", 0, false, text)
	// fmt.Printf("[Pub] end publish msg to mqtt broker, taskId: %d, count: %d, token : %s \n", taskId, i, token)
	token.Wait()

	time.Sleep(time.Duration(3) * time.Second)
	// }

	client.Disconnect(250)
}

/*
 连接任务和消息订阅方法
*/
func mqttConnSubMsgTask(taskId int) {
	// 设置连接参数
	clinetOptions := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	// 设置客户端ID
	clinetOptions.SetClientID(fmt.Sprintf("go Subscribe client example： %d-%d", taskId, time.Now().Unix()))
	// 设置连接超时
	clinetOptions.SetConnectTimeout(time.Duration(60) * time.Second)
	// 设置handler
	// clinetOptions.SetDefaultPublishHandler(messagePubHandler)
	// 创建客户端连接
	client := mqtt.NewClient(clinetOptions)

	// 客户端连接判断
	if token := client.Connect(); token.WaitTimeout(time.Duration(60)*time.Second) && token.Wait() && token.Error() != nil {
		failNums++
		fmt.Printf("[Sub] mqtt connect error, taskId: %d, fail_nums: %d, error: %s \n", taskId, failNums, token.Error())
		return
	}

	// for i := 0; i < 3; i++ {
	// 订阅消息
	token := client.Subscribe("go-test-topic", 0, messageSubHandler)
	// fmt.Printf("[Sub] end Subscribe msg to mqtt broker, taskId: %d, count: %d, token : %s \n", taskId, i, token)
	token.Wait()

	time.Sleep(time.Duration(3) * time.Second)
	// }

	client.Disconnect(250)
}
