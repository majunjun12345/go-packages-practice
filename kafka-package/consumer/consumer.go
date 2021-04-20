package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	topic     = "test"
	partition = 0 // 消费的时候需要指定 partition
	addrs     = []string{"172.20.10.3:9092"}

	wg sync.WaitGroup
)

func main() {
	singleConsumer()
}

func singleConsumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_0_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer(addrs, config)
	if err != nil {
		fmt.Printf("Failed to create sarama consumer: %v", err)
		panic(err)
	}
	defer consumer.Close()

	// 返回该消息 topic 下的所有分区
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("Failed to get partition list: %v", err)
		panic(err)
	}

	for _, partition := range partitionList {
		fmt.Printf("partition: %v\n", partition)

		// ConsumePartition 方法根据 topic，分区和给定的偏移量创建创建了相应的分区消费者
		// 如果该分区消费者已经消费了该信息将会返回error
		// sarama.OffsetNewest:表明了为最新消息，-1
		partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to consume partition: %v", err)
			panic(err)
		}
		// defer partitionConsumer.AsyncClose()

		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			// partitionConsumer.Messages() 返回消息类型的只读 chan
			for msg := range partitionConsumer.Messages() {
				fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(partitionConsumer)
	}
	wg.Wait()
}
