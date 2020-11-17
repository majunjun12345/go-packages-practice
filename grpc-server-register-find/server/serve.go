package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"testGoScripts/grpc-server-register-find/proto"
	"testGoScripts/grpc-server-register-find/register"

	"google.golang.org/grpc"
)

var (
	address string
	node    string
)

func init() {
	flag.StringVar(&address, "address", "127.0.0.1:1234", "The address")
	flag.StringVar(&node, "node", "node1", "The node")
}

type helloService struct {
}

func (h helloService) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	resp := &proto.HelloReply{} // resp := new(pb.HelloReply)
	resp.Message = "hello " + req.Name + "."
	fmt.Println("handler request", address, node)
	return resp, nil
}

func main() {
	flag.Parse()

	var (
		listen net.Listener
		err    error
	)
	fmt.Println(address)
	if err := register.InitServiceReg("hello", node, address, []string{"http://127.0.0.1:2379"}); err != nil {
		panic(err)
	}

	// 实现gRPC Server
	s := grpc.NewServer()

	// 注册helloServer为客户端提供服务
	proto.RegisterHelloServer(s, &helloService{})

	if listen, err = net.Listen("tcp", address); err != nil {
		panic(err)
	}

	s.Serve(listen)
}
