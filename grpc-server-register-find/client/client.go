package main

import (
	"context"
	"fmt"
	"sync"
	"testGoScripts/grpc-server-register-find/pool"
	pb "testGoScripts/grpc-server-register-find/proto"

	"google.golang.org/grpc"
)

var once sync.Once

func init() {
	pool.Init("http://127.0.0.1:2379")
	once.Do(func() {
		err := pool.GetConPool().Register("hello")
		if err != nil {
			panic(err)
		}
	})
}

func GetServerIns() pb.HelloClient {
	c, err := pool.GetConPool().GetClient("hello")
	if err != nil {
		panic(err)
	}
	return pb.NewHelloClient(c.(*grpc.ClientConn))
}

func main() {
	var (
		err  error
		resp *pb.HelloReply
	)

	if resp, err = GetServerIns().SayHello(context.Background(), &pb.HelloRequest{Name: "mamengli"}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Message)
}
