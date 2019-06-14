package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var cli *redis.Client

func main() {
	cli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	defer cli.Close()
	pong := cli.Ping().String() // 测试有没有连接上redis
	fmt.Println(pong)

	// StringRedis()

	HashRedis()
}

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

	cli.HScan("myinfo", 0, "", 2)
}
