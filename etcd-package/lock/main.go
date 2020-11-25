package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	v3 "github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

const (
	ServiceName = "select-master-test"
)

var (
	nodeName string
	// 参与选主的 prefix
	prefix        = "/lock"
	ETCDEndpoints = []string{"127.0.0.1:2379"}
)

func init() {
	flag.StringVar(&nodeName, "node", "node1", "Node name")
}

func main() {
	flag.Parse()
	fmt.Println("node", nodeName)

	client, err := v3.New(v3.Config{Endpoints: ETCDEndpoints})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 创建会话(租约), 会话参与选主. ttl 设定为 3 秒, 如果选主成功，会一直 keepalive, 如果 keepalive 断掉，session.Done 会收到信号
	session, err := concurrency.NewSession(client, concurrency.WithTTL(3))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// 创建 election
	election := concurrency.NewElection(session, prefix)

	// 开始竞选, 成为 master 节点的会运行起来，非 master 节点会阻塞在里面
	go func() {
		if err := election.Campaign(context.Background(), nodeName); err != nil {
			panic(err)
		}
		fmt.Println("====start work====")
	}()

	masterName := ""

	// 每隔1秒检查当前的 msster
	go func() {
		timer := time.NewTimer(time.Second * 1)
		for range timer.C {
			timer.Reset(time.Second)
			select {
			case resp := <-election.Observe(context.Background()):
				if len(resp.Kvs) > 0 {
					masterName = string(resp.Kvs[0].Value)
					fmt.Println(masterName)
				}
			}
		}
	}()

	// 每隔5s检查自己是否是 master
	go func() {
		timer := time.NewTimer(time.Second * 5)
		for range timer.C {
			timer.Reset(time.Second * 5)
			if masterName == nodeName {
				fmt.Println("master:", nodeName)
			}
		}
	}()

	select {}
}
