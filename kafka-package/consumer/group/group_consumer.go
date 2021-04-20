package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	topic     = "test"
	partition = 0 // 消费的时候需要指定 partition
	addrs     = []string{"172.20.10.3:9092"}
)

type consumerGroupHandler struct {
	groupName string
}

// Setup is run at the beginning of a new session(会话), before ConsumeClaim.
func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exites
// but before the offsets are committed for the very last time.
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("%s Message topic:%q partition:%d offset:%d  value:%s\n", h.groupName, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		// 手动确认消息
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	var (
		wg = sync.WaitGroup{}
	)
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_0_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewClient(addrs, config)
	if err != nil {
		panic(err)
	}

	g1, err := sarama.NewConsumerGroupFromClient("g1", client)
	if err != nil {
		panic(err)
	}
	defer g1.Close()
	g2, err := sarama.NewConsumerGroupFromClient("g2", client)
	if err != nil {
		panic(err)
	}
	defer g2.Close()
	g3, err := sarama.NewConsumerGroupFromClient("g3", client)
	if err != nil {
		panic(err)
	}
	defer g3.Close()
	wg.Add(3)

	// TODO: wg 这里必须是传址
	go consume(g1, &wg, "g1")
	go consume(g2, &wg, "g2")
	go consume(g3, &wg, "g3")
	wg.Wait()
}

func consume(group sarama.ConsumerGroup, wg *sync.WaitGroup, name string) {
	fmt.Println(name + " " + "start")
	defer wg.Done()
	ctx := context.Background()
	for {
		//topic := []string{"tiantian_topic1","tiantian_topic2"} 可以消费多个topic
		topics := []string{topic}
		handler := consumerGroupHandler{groupName: name}
		// TODO: group 是 interface，传进来的参数是指针，取值后能够调用 interface 内的方法
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}
