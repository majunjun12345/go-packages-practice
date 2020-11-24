package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/go-echarts/statsview"

	"testGoScripts/grpc-server-register-find/pprof"
	"testGoScripts/grpc-server-register-find/proto"
	"testGoScripts/grpc-server-register-find/register"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	reflection.Register(s)

	if listen, err = net.Listen("tcp", address); err != nil {
		panic(err)
	}

	// http://localhost:8080/debug/
	// http://localhost:8080/debug/statsviz/
	go func() {
		r := gin.New()
		pprof.Router(r)
		r.Run(":8080")
	}()

	// http://localhost:18066/debug/statsview
	// http://localhost:18066/debug/pprof/
	go func() {
		mgr := statsview.New()
		mgr.Start()
	}()

	s.Serve(listen)
}
