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
	Example()
}

func Example() {
	rConn := RedisClient.Get()
	defer rConn.Close() // 放回连接池
	resp, err := rConn.Do("PING")
	fmt.Println(resp, err)
}

func init() {
	RedisClient = RedisPool()
}

func RedisPool() *redis.Pool {
	return &redis.Pool{
		MaxActive:   50,                // 最大连接数，建议往大了配置，但不超过系统允许句柄数(ulimit -n)，0 表示没有限制
		MaxIdle:     30,                // 最大空闲连接数，会有这么多连接提前等着被使用，但超时了会被关闭
		IdleTimeout: time.Second * 300, // 空闲连接超时时间，应该设置比 redis 超时时间短，否则服务端超时了，客户端保持连接也没用
		Wait:        true,              // 如果超过最大连接数，是报错还是等待
		Dial: func() (redis.Conn, error) { // 建立连接
			conn, err := redis.Dial("tcp", RedisHost)
			if err != nil {
				return nil, err
			}
			// 密码验证
			// _, err = conn.Do("AUTH", "123456")
			// if err != nil {
			// 	conn.Close()
			// 	return nil, err
			// }
			// 选择数据库，也可以获取后再选择
			_, err = conn.Do("select", DB)
			if err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { // 没一分钟测试连接的可用状态
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
