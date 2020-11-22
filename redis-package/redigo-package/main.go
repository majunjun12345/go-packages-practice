package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

/*
	试用 redis 连接池能够有效规避每次连接的 tcp 开销

	常见报错：
		redigo: connection pool exhausted
	原因：
		当Wait==false，并且当前有效连接>=最大连接数里就报这个错了。
	解决：
		调节 MaxActive 数值
		设置 Wait 为 true
*/

const (
	RedisHost = "127.0.0.1:6379"
	DB        = 1
)

var (
	RedisClient *redis.Pool
)

func main() {
	// Example()

	// testHmsetHash()

	// testHmgetHash()

	Pipe()
}

// -----------------------hash
func testHmgetHash() {
	rConn := RedisClient.Get()
	defer rConn.Close() // 放回连接池

	data, err := redis.Values(rConn.Do("hgetall", "menglima"))
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(data); i += 2 { // 这样获取的全是 key
		fmt.Println(string(data[i].([]byte))) // 还要进行断言成 字节数组
		fmt.Println(string(data[i+1].([]byte)))
	}
}

func testHmsetHash() {
	var p1 struct {
		Title  string `redis:"title"`
		Author string `redis:"author"`
		Body   string `redis:"body"`
	}

	p1.Title = "Example"
	p1.Author = "Gary"
	p1.Body = "Hello"

	rConn := RedisClient.Get()
	defer rConn.Close() // 放回连接池

	_, err := rConn.Do("HMSET", redis.Args{}.Add("menglima").AddFlat(&p1)...)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// Pipe 管道
func Pipe() {
	conn := RedisClient.Get()
	defer conn.Close()
	conn.Send("set", "pipe1", "menglima")
	conn.Send("get", "pipe1")
	conn.Send("get", "xxx")
	conn.Flush()

	// 逐级接收结果
	r1, err := conn.Receive()
	if err != nil {
		panic(err)
	}
	fmt.Println(r1)
	r2, err := conn.Receive()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(r2.([]byte)))
	r3, err := conn.Receive()
	if err != nil {
		panic(err)
	}
	fmt.Println(r3)
}

// Example pool
func Example() {
	rConn := RedisClient.Get()
	defer rConn.Close() // 放回连接池

	// err 这样就把连接池关了
	// rConn1 := RedisClient
	// defer rConn1.Close()
	// rConn1.Get() ...

	resp, err := rConn.Do("PING")
	fmt.Println(resp, err)
}

func init() {
	RedisClient = RedisPool()
}

// RedisPool redis conn pool
func RedisPool() *redis.Pool {
	return &redis.Pool{
		MaxActive:   50,                // 最大的激活连接数，表示同时最多有N个连接
		MaxIdle:     30,                // 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态，但超时了会被关闭
		IdleTimeout: time.Second * 300, // 空闲连接超时时间，应该设置比 redis 超时时间短，否则服务端超时了，客户端保持连接也没用
		// Wait:        true,              // 如果超过最大连接数，是报错还是等待
		Dial: func() (redis.Conn, error) { // 建立连接
			conn, err := redis.Dial("tcp", RedisHost)
			if err != nil {
				return nil, err
			}
			// 密码验证
			_, err = conn.Do("AUTH", "123456")
			if err != nil {
				conn.Close()
				return nil, err
			}
			// 选择数据库，也可以获取后再选择
			_, err = conn.Do("select", DB)
			if err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { // 每一分钟测试连接的可用状态
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
