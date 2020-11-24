package main

import (
	"context"
	"fmt"
	"sync"
	"testGoScripts/grpc-server-register-find/pool"
	pb "testGoScripts/grpc-server-register-find/proto"
	"testGoScripts/grpc-server-register-find/tool"
	"testGoScripts/grpc-server-register-find/tool/tracer"
	"testGoScripts/zaplog"
	"time"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

var once sync.Once

func init() {
	pool.Init("http://127.0.0.1:2379")
	once.Do(func() {
		err := pool.GetConPool().FindServer("hello")
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
	tool.InitLog()
	tracer.InitGlobal("hello", zaplog.GetLogger())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
	defer cancel()

	if span := opentracing.SpanFromContext(ctx); span != nil {
		newSpan := opentracing.StartSpan("SayHello", opentracing.ChildOf(span.Context()))
		defer newSpan.Finish()
		newSpan.SetTag("type", "SayHello client type")
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
	}

	if resp, err = GetServerIns().SayHello(ctx, &pb.HelloRequest{Name: "mamengli"}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Message)
}
