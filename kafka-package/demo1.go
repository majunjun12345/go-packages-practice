package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

/*
	ubuntu 启动 kafka 服务端：
	bin/zookeeper-server-start.sh config/zookeeper.properties
	bin/kafka-server-start.sh config/server.properties

	提供外网服务：
	vim config/server.properties 修改：
	取消注释：listeners=PLAINTEXT://:9092
	填写本机 ip(不能为 0.0.0.0)：advertised.listeners=PLAINTEXT://172.31.144.70:9092

	本例启动方式：
	command 识别是消费端还是生产端

	github.com/bsm/sarama-cluster 结合 github.com/Shopify/sarama 能够追踪 offset

	kafka 集群的 demo：
	https://github.com/ErikJiang/kafka_cluster_example/wiki/GoKafkaPubSubBuild
*/

var (
	topic     = "test"
	partition = 0 // 消费的时候需要指定 partition
	addrs     = []string{"192.168.0.103:9092"}

	saslEnable = false
	tlsEnable  = false
	clientcert = ""
	clientkey  = ""
	cacert     = ""

	// consumer producer
	command = "consumer"
)

func main() {
	// kafka()
	metadata()
}

// 通过 client 创建 consumer 或 producer
func kafka() {
	conf := sarama.NewConfig()
	// 等待服务器将所有副本都保存成功后的响应
	conf.Producer.RequiredAcks = sarama.WaitForAll
	// 随机向 partition 发送消息：返回一个分区器，该分区器每次选择一个随机分区
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功或失败后的响应，只有上面的RequireAcks设置不是NoReponse这里才有用
	conf.Producer.Return.Successes = true

	// 安全认证方面的参数设置
	if saslEnable {
		conf.Net.SASL.Enable = true
		conf.Net.SASL.User = "username"
		conf.Net.SASL.Password = "password"
	}

	// https
	if tlsEnable {
		//sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
		tlsConfig, err := genTLSConfig(clientcert, clientkey, cacert)
		if err != nil {
			log.Fatal(err)
		}

		conf.Net.TLS.Enable = true
		conf.Net.TLS.Config = tlsConfig
	}

	client, err := sarama.NewClient(addrs, conf)
	if err != nil {
		fmt.Printf("Failed to create sarama client: %v\n", err)
	}

	if command == "producer" {
		// 并发量较小时可以用 NewSyncProducerFromClient
		producer, err := sarama.NewAsyncProducerFromClient(client)
		if err != nil {
			log.Fatal(err)
		}
		defer producer.Close()
		loopProducer(producer, topic)
	}
	if command == "consumer" {
		consumer, err := sarama.NewConsumerFromClient(client)
		if err != nil {
			log.Fatal(err)
		}
		defer consumer.Close()
		loopConsumer(consumer, topic, partition)
	}
}

func genTLSConfig(clientcertfile, clientkeyfile, cacertfile string) (*tls.Config, error) {
	// load client cert
	clientcert, err := tls.LoadX509KeyPair(clientcertfile, clientkeyfile)
	if err != nil {
		return nil, err
	}

	// load ca cert pool
	cacert, err := ioutil.ReadFile(cacertfile)
	if err != nil {
		return nil, err
	}
	cacertpool := x509.NewCertPool()
	cacertpool.AppendCertsFromPEM(cacert)

	// generate tlcconfig
	tlsConfig := tls.Config{}
	tlsConfig.RootCAs = cacertpool
	tlsConfig.Certificates = []tls.Certificate{clientcert}
	tlsConfig.BuildNameToCertificate()
	// tlsConfig.InsecureSkipVerify = true // This can be used on test server if domain does not match cert:
	return &tlsConfig, err
}

func loopConsumer(consumer sarama.Consumer, topic string, partition int) {
	partitionConsumer, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
	if err != nil {
		log.Println(err)
		return
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message: [%s], offset: [%d]\n", msg.Value, msg.Offset)
		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}

	}
}

func loopProducer(producer sarama.AsyncProducer, topic string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		if text == "exit" || text == "quit" {
			break
		}

		producer.Input() <- &sarama.ProducerMessage{
			Topic: topic,
			Key:   nil,
			Value: sarama.StringEncoder(text),
		}
		log.Printf("Produced message: [%s]\n", text)

		// 判断消息是否生产成功
		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %d\n", suc.Offset, suc.Timestamp.Unix())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
		fmt.Print("> ")
	}
}

// metadata 元数据
func metadata() {
	var (
		config  = sarama.NewConfig()
		client  sarama.Client
		topics  []string
		brokers []*sarama.Broker
		err     error
	)

	if client, err = sarama.NewClient(addrs, config); err != nil {
		fmt.Printf("metadata try create client err :%s\n", err.Error())
		return
	}
	defer client.Close()

	// get topic set
	if topics, err = client.Topics(); err != nil {
		fmt.Printf("try get topics err %s\n", err.Error())
		return
	}
	fmt.Printf("topics(%d):\n", len(topics))
	for _, topic := range topics {
		fmt.Println(topic)
	}

	brokers = client.Brokers()
	fmt.Printf("broker set(%d):\n", len(brokers))
	for _, broker := range brokers {
		fmt.Printf("%s\n", broker.Addr())
	}

	partion, err := client.Partitions(topic)
	if err != nil {
		panic(err)
	}
	fmt.Println("partion:", partion)
}
