docker run -d --name kafka -p 9092:9092 -e KAFKA_BROKER_ID=0 -e KAFKA_ZOOKEEPER_CONNECT=docker.for.mac.host.internal:2181/kafka -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://192.168.0.103:9092 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 wurstmeister/kafka

KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://192.168.0.103:9092，为本机实际 ip，不能为 0.0.0.0

docker run -d --name zookeeper -p 2181:2181 wurstmeister/zookeeper

[再过半小时，你就能明白kafka的工作原理了](https://www.cnblogs.com/sujing/p/10960832.html)
[震惊了！原来这才是kafka！](https://www.jianshu.com/p/d3e963ff8b70)
[一文彻底搞懂Kafka](https://mp.weixin.qq.com/s/c45nQI9nLT5lwu7vSvHgjg)

### kafka 简介

#### 一、mq 的几种消息传递方式
- 点对点模式
  消息持久化到一个topic中，此时会有一个或多个消费topic中的消息，但是一条消息只能被消费一次，被消费后则从队列中删除该消息

  ![](https://images2018.cnblogs.com/blog/1228818/201805/1228818-20180507190326476-771565746.png)

    **生产者发送一条消息到queue，只有一个消费者能收到**

- 发布订阅模式
  消息持久化到一个topic中，与点对点不同的是，消费者可以订阅一个或多个topic，消费者可以消费该topic中所有的数据，同一条数据可以被多个消费者消费，数据被消费后不会立马删除。
  
  ![](https://images2018.cnblogs.com/blog/1228818/201805/1228818-20180507190443404-1266011458.png)

  **发布者发送到topic的消息，订阅了topic的订阅者都会收到消息。**

#### 二、kafka 是什么
发布与订阅消息
作为消息中间件使用
实时流数据传输和处理
kafka 的 HA 是通过 数据复制和leader选举来保证的
时间复杂度 O(1)

#### 三、应用场景
  日志收集系统和消息系统。

#### 四、kafka 相关名次解释
- producer
  消息和数据的生产者，向 kafka 的一个 topic 发布消息服务；
- consumer
  消息和数据的消费者，订阅数据(topic)并处理其发布的消息；
- consumer group
  对于同一个 topic，会广播给不同的 group，只有一个 consumer 可以消费该消息；
- broker
  Kafka 集群包含一个或多个服务器，服务器节点称为broker；
- topic
  kafka 消息的类别，对数据(消息)进行分区、隔离
- partition
  kafka 下数据存储的节本单元，一个 topic 数据会被分散存储到一个或多个 partition中，每一个 partition 内数据是有序的，但是不能保证不同 partition 之间消息的顺序。
  是消费者消费的基本单位，一个 partition 至多被一个消费者消费。单 partition 被多消费者消费可以使用消费组；原因：为保证顺序，kafka 强制点对点按顺序传递消息，意味着 consumer 在分区中的位置只有一个；
    - 几个原则
      1. partition在写入的时候可以指定需要写入的partition，如果有指定，则写入对应的partition。
      2. 如果没有指定partition，但是设置了数据的key，则会根据key的值hash出一个partition。
      3. 如果既没指定partition，又没有设置key，则会采用轮询⽅式，即每次取一小段时间的数据写入某个partition，下一小段的时间写入下一个partition
    - 结构
      每个 partition 的文件夹下会有多组 segment 文件夹
      每个 segment 文件夹都包含 .index、 .log、 .timeindex 三个文件
      .log ⽂件存储 message，.index 和 .timeindex 为索引⽂件，⽤于检索消息。
- replication(副本)
  同一个 partition 可能会有多个副本，多个副本之间的数据是一样的；
- 提交
  更新 partition 当前位置的操作叫提交；

![image.png](https://upload-images.jianshu.io/upload_images/7998142-e48724b8c47d7be8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

大概用法就是，Producers往Brokers里面的指定Topic中写消息，Consumers从Brokers里面拉去指定Topic的消息，然后进行业务处理。

一个 topic 由一个或多个 partition 组成，Kafka无法在整个topic范围内保证消息的顺序，但可以保证消息在单个分区内的顺序。

Consumers 数量小于等于 partition 数量，避免同一个 partition 被多个 sonsumer 消费；

partition 在物理上对应一个文件夹，存储 partition 的所有消息数据(hashkey)和索引文件，当遇到瓶颈时，可以通过增加 partition 的数量来横向扩容；

图中有两个topic，topic 0有两个partition，topic 1有一个partition，每个 partition 都会有多个副本。当存在多副本的情况下，尽量会把副本分配到不同的 broker 上。kafka 会为 partition 选出一个 leader，之后所有该 partition 的请求实际操作的都是 leader，然后同步到其他 follower；

- 数据的清理
对于 kafka 来说，磁盘是最重要的子系统，所有的消息都保存在磁盘中，所以 kafka 的性能严重依赖磁盘的性能；

- kafka 为什么吞吐量大
Kafka大量使用操作系统页缓存，内存操作速度快且命中率高
Kafka不直接参与物理I/O操作，而是交由最擅长此事的操作系统来完成
顺序存储和MMFile
Kafka使用顺序I/O进行读写
使用以sendfile为代表的零拷贝技术加强了网络间的数据传输效率

- ACK应答机制
  0 代表producer往集群发送数据不需要等到集群的返回，不确保消息发送成功。安全性最低但是效 率最高。
  1 代表producer往集群发送数据只要leader应答就可以发送下一条，只确保leader发送成功。
  all 代表producer往集群发送数据需要所有的follower都完成从leader的同步才会发送下一条，确保 leader发送成功和所有的副本都完成备份。安全性最⾼高，但是效率最低。

#### 五、kafka 的消费
- 没有消费组的消费者
  kafka 的消费必须指定 partition，如果该消费者没有消费组，那么其只能消费指定 partition 的消息;
- 消费组内的消费者
  consumer group是kafka提供的可扩展且具有容错性的消费者机制。既然是一个组，那么组内必然可以有多个消费者或消费者实例(consumer instance)，它们共享一个公共的ID，即group ID。组内的所有消费者协调在一起来消费订阅主题(subscribed topics)的所有分区(partition)。当然，每个分区只能由同一个消费组内的一个consumer来消费。个人认为，理解consumer group记住下面这三个特性就好了：
  - consumer group下可以有一个或多个consumer instance，consumer instance可以是一个进程，也可以是一个线程
  - group.id是一个字符串，唯一标识一个consumer group
  - consumer group下订阅的topic下的每个分区只能分配给某个group下的一个consumer(当然该分区还可以被分配给其他group)
- 消费组的 **`rebalance`**
  rebalance规定了一个consumer group下的所有consumer如何达成一致来分配订阅topic的每个分区。比如某个group下有20个consumer，它订阅了一个具有100个分区的topic。正常情况下，Kafka平均会为每个consumer分配5个分区。这个分配的过程就叫rebalance。
  rebalance的触发条件有三种：
  - 组成员发生变更(新consumer加入组、已有consumer主动离开组或已有consumer崩溃了——这两者的区别后面会谈到)
  - 订阅主题数发生变更——这当然是可能的，如果你使用了正则表达式的方式进行订阅，那么新建匹配正则表达式的topic就会触发rebalance
  - 订阅主题的分区数发生变更
- 消费状态
  kafka 不保存消息为状态，即是否被消费；一般的消息系统需要保存消息的状态，并且还需要以随机访问的形式更新消息的状态。而Kafka 的做法是保存Consumer在Topic分区中的位置offset，在offset之前的消息是已被“消费”的，在offset之后则为未“消费”的，并且offset是可以任意移动的，这样可以消除了大部分的随机IO。

- 延时消费
##### 六、注意事项
[go 操作 kafka](https://www.cnblogs.com/gwyy/p/13266589.html)

普通消费者（我姑且这么说）。有些情况下我们有些消费者是没有消费组的，正常的消费者可自动分配分区到消费者并且组中消费者新增或删除会自动触**发负载均衡的消费组**。

但在某些情况下，却想要更简单的东西。有时你知道你有一个单一的消费者总是需要从主题中的所有分区读取数据，或者从一个主题特定分区读取数据。
在这种情况下没有理由需要组或负载均衡，只是订阅特定的主题或分区，偶尔使用消息和提交偏移量。

但是有个注意的点。除了没有负载均衡以及需要手动查找分区，一切看起来都很正常。请记住，如果有人向主题添加新分区，则不会通知消费者。所以无论是处理通过定期检查consumer.partitionsFor()或者记住是否是管理员添加分区，应用程序将需要跳跃。还要注意的是消费者可以订阅的主题（成为一个消费组的一部分），或分配自己的分区，但不能同时实现。下面可以看看代码。一般不这么用。一般都用消费组+消费者
```go
func main() {
   var wg sync.WaitGroup
   //创建消费者
   config := sarama.NewConfig()
   config.Consumer.Return.Errors = true
   client,err := sarama.NewClient([]string{"10.180.18.60:9092"}, config)
   defer client.Close()
   if err != nil {
      panic(err)
   }
   consumer, err := sarama.NewConsumerFromClient(client)
   defer consumer.Close()
   if err != nil {panic(err)}
   //设置分区
   partitionList, err :=  consumer.Partitions("liangtian_topic")
   if err != nil {
      fmt.Println("faild to get the list of partitions",err)
   }
   //[0 1 2]
   fmt.Println(partitionList)
   //循环读取分区
   for partition := range partitionList {
      pc, err := consumer.ConsumePartition("liangtian_topic", int32(partition), sarama.OffsetOldest)
      if err != nil {
         fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
         return
      }
      defer pc.AsyncClose()
      wg.Add(1)
      go func(pc sarama.PartitionConsumer) {
         defer wg.Done()
         for msg := range pc.Messages() {
            fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
            fmt.Println()
         }
      }(pc)
   }
   //time.Sleep(time.Hour)
   wg.Wait()
   consumer.Close()
}
```