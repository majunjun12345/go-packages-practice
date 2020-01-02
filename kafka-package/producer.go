package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/golang/glog"
)

func main() {
	// syncProducer()
	// consumer()
	metadata()
}

func init() {

}

// 消息量大必须用异步生产
func asyncProducer() {
	config := sarama.NewConfig()
	config.Producer.Timeout = 5 * time.Second
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
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case err := <-p.Errors():
				if err != nil {
					glog.Errorln(err)
				}
			case suc := <-p.Successes():
				fmt.Printf("producer success: %v\n", suc.Offset)
			}
		}
	}(producer)

	msg := &sarama.ProducerMessage{
		Topic: topic,
		//Key:   sarama.StringEncoder("go_test"),
		Value: sarama.ByteEncoder("this is message"),
	}
	producer.Input() <- msg
	select {}
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

	for i := 0; i < 5; i++ {
		//创建消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = topic
		msg.Value = sarama.StringEncoder("this is a good test,hello kai")
		//发送消息
		pid, offset, err := syncProducer.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		time.Sleep(time.Second)
	}
}

func consumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_0_0_0

	// consumer
	consumer, err := sarama.NewConsumer(addrs, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}
	defer consumer.Close()

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("consumer get partitions error %s\n", err)
	}

	for _, partition := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("try create partition_consumer error %s\n", err.Error())
			continue
		}

		for {
			select {
			case msg := <-partitionConsumer.Messages():
				fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
					msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
			case err := <-partitionConsumer.Errors():
				fmt.Printf("err :%s\n", err.Error())
			}
		}
	}
}

// 元数据
func metadata() {
	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_2

	client, err := sarama.NewClient(addrs, config)
	if err != nil {
		fmt.Printf("metadata_test try create client err :%s\n", err.Error())
		return
	}

	defer client.Close()

	// get topic set
	topics, err := client.Topics()
	if err != nil {
		fmt.Printf("try get topics err %s\n", err.Error())
		return
	}

	fmt.Printf("topics(%d):\n", len(topics))

	for _, topic := range topics {
		fmt.Println(topic)
	}

	// get broker set
	brokers := client.Brokers()
	fmt.Printf("broker set(%d):\n", len(brokers))
	for _, broker := range brokers {
		fmt.Printf("%s\n", broker.Addr())
	}
}
