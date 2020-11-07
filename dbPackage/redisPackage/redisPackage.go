package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var cli *redis.Client

func main() {
	cli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       1,
	})
	defer cli.Close()
	pong := cli.Ping().String() // 测试有没有连接上redis
	fmt.Println(pong)

	// StringRedis()

	// HashRedis()

	// ListRedis()

	// SetRedis()

	// ZsetRedis()

	// PipelineRedis()

	// TransactionRedis()

	// PubSubRedis()

	// Hset()

	HashRedis()
}

func Hset() {
	info := make(map[string]interface{})
	info["status"] = "online"
	// info["online_time"] = 1
	info["update_time"] = time.Now().UnixNano() / 1e6
	err := cli.HMSet("device_id", info).Err()
	if err != nil {
		panic(err)
	}
	time.Sleep(3 * time.Second)
	err = cli.HIncrBy("device_id", "online_time", 5).Err()
	if err != nil {
		panic(err)
	}

	// err = cli.HSet("device_id", "online_time", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// info2 := make(map[string]interface{})
	// info2["status"] = "offline"
	// // info["online_time"] = 1
	// info2["update_time"] = time.Now().UnixNano() / 1e6
	// err = cli.HMSet("device_id", info2).Err()
	// if err != nil {
	// 	panic(err)
	// }
}

// hash
func HashRedis() {
	cli.HSet("myinfo", "name", "menglima").Result()
	info := make(map[string]interface{})
	info["age"] = 19
	info["home"] = "xiantao"
	cli.HMSet("myinfo", info)
	// fmt.Println(cli.HExists("myinfo", "name").Result())
	// fmt.Println(cli.HDel("myinfo", "name").Result())
	// fmt.Println(cli.HExists("myinfo", "name").Result())
	fmt.Println(cli.HGetAll("myinfo").Result())
	fmt.Println(cli.HIncrBy("myinfo", "age", 1).Result())
	fmt.Println(cli.HKeys("myinfo").Result()) // keys 列表
	fmt.Println(cli.HVals("myinfo").Result()) // values 列表
	fmt.Println(cli.HLen("myinfo").Result())
	fmt.Println(cli.HSetNX("myinfo", "phone", "189").Result())

	// iter := cli.HScan("myinfo", 0, "", 1).Iterator() // 参数:redis的键,从哪里开始迭代, 匹配的value, 每次获取 value的个数 (hash 的 key value 是分开的)
	// for iter.Next() {
	// 	fmt.Println("1111")
	// 	iter.Next() // 这样可以忽略 hash 的 key
	// 	fmt.Println("iter value:", iter.Val())
	// }

	// results, _ := cli.Scan(0, "device_", 100).Val()
	n, _ := cli.HIncrBy("device_id_01", "1", 5).Result()
	fmt.Println("======", n)
}

// string
func StringRedis() {
	// 0 表示永久有效
	fmt.Println(cli.Set("key", "101", 0).Args()) // [set key value]
	// fmt.Println(cli.Set("key", "value", 0).Err())    // <nil>
	// fmt.Println(cli.Set("key", "value", 0).Name())   // set
	// fmt.Println(cli.Set("key", "value", 0).String()) // set key value:OK
	// fmt.Println(cli.Set("key", "value", 0).Val())    // OK
	// fmt.Println(cli.Set("key", "value", 0).Result()) // OK <nil>

	// get
	result, err := cli.Get("2wferf").Result() // hahaha <nil>
	if err == redis.Nil {                     // redis.Nil 来判断有没有该 key
		fmt.Println("no such key")
	}
	fmt.Println(result)
	// fmt.Println(cli.Get("key").Args())        // [get key]
	// fmt.Println(cli.Get("key").Bytes())       // [104 97 104 97 104 97] <nil>
	// fmt.Println(cli.Get("key").Err())         // <nil>
	// fmt.Println(cli.Get("key").Float32())     // 直接将值转换为数值, 字符串数值可以转,字符串字符不能转
	// fmt.Println(cli.Get("key").Float64())     //
	// fmt.Println(cli.Get("key").Int())         //
	// fmt.Println(cli.Get("key").Int64())       //
	// fmt.Println(cli.Get("key").Uint64())      //
	// fmt.Println(cli.Get("key").Name())        // get
	// fmt.Println(cli.Get("key").String())      // get key:hahaha
	// fmt.Println(cli.Get("key1").Val())        // 如果没有则为空, 一般是通过下面的方式获取值

	// getset
	res, err := cli.GetSet("key", 102).Result() // 和 setNX 不一样
	if err != redis.Nil {
		fmt.Println(res)
	}

	fmt.Println(cli.MGet("key", "name").Result()) // [102 majun]

	fmt.Println(cli.SetNX("key", "103", 0).Result())  // 当 key 存在时,不设置,返回 false
	fmt.Println(cli.SetNX("key1", "103", 0).Result()) // key 不存在时,设置,返回 false
	cli.Incr("key1").Result()                         // 自增 1  数值才有效
	cli.DecrBy("key1", 4).Result()                    // 自减 n
	cli.Append("key1", "sanqi").Result()              //  相当于字符串的 +
}

// list
func ListRedis() {

	// cli.LPush("lnames", "masanqi", "1", "2", "3", "4") // 新建并插入元素, 可以一次性插入多个, 最后面参数在redis最前面
	// fmt.Println(cli.LPushX("lnames", "menglima").Result()) // 插入到头部
	// cli.RPush("lnames", "lisi") // 插入到尾部
	// fmt.Println(cli.LIndex("lnames", -1).Result()) // 获取指定索引的值
	// fmt.Println(cli.LRange("lnames", 0, -1).Result()) // 获取 index 区间值
	// fmt.Println(cli.LLen("lnames").Result()) // 获取 list 长度

	// fmt.Println(cli.LSet("lnames", 1, "l1")) // 在指定 index 插入
	// cli.LInsert("lnames", "BEFORE", "lisi", "wangwu") // 在指定值 前/后 插入
	// cli.LInsertAfter("lnames", "lisi", "sunliu")
	// cli.LInsertBefore("lnames", "menglima", "hahaha") // 在发现的第一个元素前插入

	// cli.LPop("lnames").Result() // 删除头部
	// cli.RPop("lnames") // 删除尾部
	// fmt.Println(cli.LRem("lnames", 2, "menglima").Result()) // 删除指定元素个数
	// cli.LTrim("lnames", -2, -1) // 只保留指定区间的值

	cli.RPopLPush("lnames", "lnumbers") // 将左尾移到右头
}

// set
func SetRedis() {

	// 无序集合,元素不能重复
	// cli.SAdd("sid", "01", "02", "03", "04", "05", "06") // 添加元素,没有则创建
	// cli.SAdd("sid1", "04", "05", "06", "07", "08")
	// fmt.Println(cli.SCard("sid"))   // 获取元素个数

	// fmt.Println(cli.SDiff("sid", "sid1")) // 获取两个 key 所对应的集合的差集
	// fmt.Println(cli.SDiffStore("sid2", "sid", "sid1")) // sid 和 sid1 的差集存储到 sid2 中
	// cli.SInter()  // 交集
	// cli.SUnion()  // keys 的并集

	// fmt.Println(cli.SIsMember("sid", "01").Result()) // 是否存在该元素
	// fmt.Println(cli.SRandMember("sid").Result()) //随机获取几个或多个元素
	// fmt.Println(cli.SMembers("sid")) // 获取所有元素

	// cli.SAdd("sid3")
	// cli.SMove("sid", "sid3", "04") // 将指定元素从 sid 移到 sid3

	// fmt.Println(cli.SPop("sid").Result())     // 移除随机元素
	// fmt.Println(cli.SPopN("sid", 2).Result()) // 移除随机元素
	// fmt.Println(cli.SRem("sid", "07").Result()) // 移除指定元素

	iter := cli.SScan("sid", 0, "", 1).Iterator()
	for iter.Next() {
		fmt.Println(iter.Val())
	}
}

func ZsetRedis() {
	// scores := []*redis.Z{}
	// score1 := &redis.Z{98, "math"}
	// score2 := &redis.Z{89, "eng"}
	// score3 := &redis.Z{76, "chi"}
	// score4 := &redis.Z{63, "phy"}
	// score5 := &redis.Z{52, "che"}
	// score6 := &redis.Z{71, "geo"}
	// scores = append(scores, score1, score2, score3, score4, score5, score6)
	// fmt.Println(cli.ZAdd("zscore", scores...).String())

	/*
		Zadd 更新或插入
		ZAddNX 不改变已有
		ZAddXX 只改变已有的
		ZAddch 返回被修改的元素的个数,返回 0 表示没有被修改.返回 1 表示被修改或者插入;
	*/
	// fmt.Println(cli.ZAdd("zscore", &redis.Z{101, "eng"}).String())
	// fmt.Println(cli.ZAddCh("zscore", &redis.Z{101, "eng"}).String())
	// fmt.Println(cli.ZAddXX("zscore", &redis.Z{89, "eng"}).String())
	// fmt.Println(cli.ZAddCh("zscore", &redis.Z{89, "xxx"}).Result())

	// fmt.Println(cli.ZCard("zscore"))                       // zset 中的成员数
	// fmt.Println(cli.ZCount("zscore", "0", "100").Result()) // 分数区间的元素个数
	// fmt.Println(cli.ZRange("zscore", 0, 100).Result()) // 返回介于分数区间的元素
	// fmt.Println(cli.ZRangeWithScores("zscore", 0, 100).Result()) // 返回介于分数区间的元素及对应的分数
	// fmt.Println(cli.ZRangeByScore("zscore", &redis.ZRangeBy{"0", "100", 0, 0}).Result()) // 返回分数区间对应的元素,还有参数 offset 和 count

	// 有坑, 暂时还不会用.原因是 cli 获取到的是 二进制, 好像不能直接映射为 go 的对象
	// results := make([]MyStruct, 1)
	// fmt.Println(cli.ZRangeByScore("zscore", &redis.ZRangeBy{"0", "100", 1, 1}).ScanSlice(&results)) // 返回分数区间对应的元素,还有参数 offset 和 count
	// fmt.Printf("%+v", results)

	// fmt.Println(cli.ZRank("zscore", "===").String()) // 获取元素排名, 从低到高
	// fmt.Println(cli.ZRem("zscore", "===").Result()) // 删除元素
	// fmt.Println(cli.ZRevRange("zscore", 0, 100).Result()) // 区间内所有元素,从大到小排列
	// fmt.Println(cli.ZScore("zscore", "eng").Result()) // 获取指定元素的分数

	/*
		ZIncrNX: 没有就插入
		ZIncrXX: 只改变已有的
	*/
	// fmt.Println(cli.ZIncrBy("zscore", 5, "eng").Result())
	// fmt.Println(cli.ZIncrNX("zscore", &redis.Z{5, "==="}).Result())
	// fmt.Println(cli.ZIncrXX("zscore", &redis.Z{5, "[{{{{"}).Result())

	// cli.ZInterStore("zscore1", &redis.ZStore{}, "zscore").Result() // 将 keys 的交集存储在新的 key 中

	// fmt.Println(cli.ZLexCount("zsocre", "[80", "[100").Result())   // 不知道是干啥的

	// iter := cli.ZScan("zscore", 0, "", 1).Iterator()
	// for iter.Next() {
	// 	iter.Next() // 忽略 元素,只获取分数值
	// 	fmt.Println(iter.Val())
	// }

}

type MyStruct redis.Z

func (m *MyStruct) UnmarshalBinary(data []byte) error {
	// convert data to yours, let's assume its json data
	return json.Unmarshal(data, m)
	// encoding.BinaryUnmarshaler()
	// u := url.URL{}
	// u.UnmarshalBinary()
}

/*
	管道
	管道可以理解为一系列命令的打包. 通常, redis 的 cli 与 server 交互时,都是一个命令执行完后,明确收到 server 的反馈信息后才执行下一个命令, 这种交互是堵塞式的,效率比较低下.
	使用管道后, 多个命令可以放到一起执行, 其实在管道中 redis 还是依次执行每个命令, 不过下个命令不用等到上个命令的执行结果反馈到 client 后再执行;
*/

func PipelineRedis() {
	t1 := time.Now().Unix()
	for i := 0; i < 100000000; i++ {
		p := cli.Pipeline()
		p.Set("num", 0, 0)
		p.Incr("num")
		p.Incr("num")
		p.Incr("num")
		_, err := p.Exec()
		if err != nil {
			panic(err)
		}
	}
	t2 := time.Now().Unix()
	fmt.Println(t2 - t1)

	fmt.Println(cli.Get("num").Result())
}

func NoPipelineRedis() {
	t1 := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		cli.Set("num", 0, 0)
		cli.Incr("num")
		cli.Incr("num")
		cli.Incr("num")
	}
	t2 := time.Now().Unix()
	fmt.Println(t2 - t1)

	fmt.Println(cli.Get("num").Result())
}

/*
	通过 multi 开启事物, 通过 exec 提交事物, discard 回滚一个操作, watch 和 unwatch 可以监控和取消监控指定的 key.
	事物与管道类似, 也是将多个命令打包, 然后放到一起执行, 区别是如果有命令执行失败, 则回滚;
*/
func TransactionRedis() {
	tx := cli.TxPipeline()
	tx.Set("num", 0, 0)
	tx.Incr("num")
	tx.Incr("num")
	tx.Incr("num")
	_, err := tx.Exec()
	if err != nil {
		panic(err)
	}
}

// pub/sub
/*
	pub 和 sub 必须在两个不同的 cli 中使用
*/
func PubSubRedis() {
	fmt.Println(cli.Publish("mychannel", "hello redis!").Err())
	// sub := cli.Subscribe("mychannel")
	// msgChan := sub.Channel()
	// for {
	// 	select {
	// 	case msg := <-msgChan:
	// 		fmt.Println(msg.String())
	// 	}
	// }
}
