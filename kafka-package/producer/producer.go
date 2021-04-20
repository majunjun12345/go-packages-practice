package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/golang/glog"
	"github.com/spf13/cast"
)

var (
	topic     = "test"
	partition = 0 // 消费的时候需要指定 partition
	addrs     = []string{"172.20.10.3:9092"}

	saslEnable = false
	tlsEnable  = false
	clientcert = ""
	clientkey  = ""
	cacert     = ""
)

func main() {
	// syncProducer()
	asyncProducer()
}

// 消息量大必须用异步生产
func asyncProducer() {
	config := sarama.NewConfig()
	config.Producer.Timeout = 5 * time.Second
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机分区
	config.Producer.RequiredAcks = sarama.WaitForAll          //等待服务器所有副本都保存成功后的响应
	config.Producer.Partitioner = sarama.NewRandomPartitioner //随机向partition发送消息
	config.Producer.Return.Successes = true                   //是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用. 必须有这个选项
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
	config.Version = sarama.V2_0_0_0

	producer, err := sarama.NewAsyncProducer(addrs, config)
	if err != nil {
		fmt.Printf("producer_test create producer error :%s\n", err.Error())
		return
	}
	defer producer.AsyncClose()

	// 创建协程用于接收异步生产结果通知，必须是在协程里面，不然阻塞主进程
	// 如果打开了Return.Successes和Errors配置，而又没有p.Successes()提取，那么Successes()这个chan消息会被写满导致死锁。
	go func() {
		for {
			select {
			case err := <-producer.Errors():
				if err != nil {
					glog.Errorln(err)
				}
			case suc := <-producer.Successes():
				fmt.Printf("producer success: %v\n", suc.Offset)
			}
		}
	}()

	client, err := sarama.NewClient(addrs, config)
	if err != nil {
		panic(err)
	}

	for {

		msg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder("go_test"),
			Value: sarama.ByteEncoder("this is message"),
		}
		// 异步发送消息的方式
		producer.Input() <- msg

		p, err := client.Partitions(topic)
		if err != nil {
			panic(err)
		}
		fmt.Println("======partitions:", p)

		time.Sleep(time.Second)
	}
}

// 同步生产
func syncProducer() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	syncProducer, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer close,err:", err)
		return
	}
	defer syncProducer.Close()

	for i := 0; i < 100000000; i++ {
		//创建消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = topic
		msg.Value = sarama.StringEncoder("this is a good test,hello kai" + " " + cast.ToString(i))
		// 同步发送消息的方式
		partition, offset, err := syncProducer.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}
		fmt.Printf("partition:%v offset:%v\n", partition, offset)
		time.Sleep(time.Second)
	}
}
