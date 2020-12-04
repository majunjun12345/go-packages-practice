package main

import (
	"context"
	"flag"
	"net"
	"time"

	"github.com/go-echarts/statsview"
	"github.com/opentracing/opentracing-go"

	"testGoScripts/grpc-server-register-find/pprof"
	"testGoScripts/grpc-server-register-find/proto"
	"testGoScripts/grpc-server-register-find/register"
	"testGoScripts/grpc-server-register-find/tool"
	"testGoScripts/grpc-server-register-find/tool/tracer"
	"testGoScripts/zaplog"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

const (
	SERVER_NAME = "hello"
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
	return &proto.HelloReply{Message: "hello " + req.Name + "."}, nil
}

func main() {
	flag.Parse()
	tool.InitLog()

	tracer.InitGlobal(SERVER_NAME, zaplog.GetLogger())

	var (
		listen net.Listener
		err    error
	)
	if err := register.InitServiceReg(SERVER_NAME, node, address, []string{"http://127.0.0.1:2379"}); err != nil {
		panic(err)
	}

	// 实现gRPC Server
	s := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             time.Second * 5, // 如果客户端 ping 的频率超过 每次/5s, 则断开该连接
			PermitWithoutStream: true,            // 果没有active的stream,是否允许发送ping
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     time.Second * 15, // 如果客户端超过 15s 没有通信，则断开连接
			MaxConnectionAge:      time.Second * 30, // 连接的最长时间
			MaxConnectionAgeGrace: time.Second * 5,  // 强制断开的缓冲时间 MaxConnectionAge + 5
			Time:                  time.Second * 10, // 每 10s ping 一次客户端，必须在 1s 内返回，否则回收连接
			Timeout:               time.Second,
		}),
		grpc.UnaryInterceptor(tracer.OpentracingServerInterceptor(opentracing.GlobalTracer())),
	)

	// 注册helloServer为客户端提供服务
	proto.RegisterHelloServer(s, &helloService{})

	// grpcui
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
