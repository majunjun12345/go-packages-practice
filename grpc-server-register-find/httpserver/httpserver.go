package main

import (
	"context"
	"net/http"
	"sync"
	"testGoScripts/grpc-server-register-find/pool"
	pb "testGoScripts/grpc-server-register-find/proto"
	"testGoScripts/grpc-server-register-find/tool"
	"testGoScripts/grpc-server-register-find/tool/tracer"
	"testGoScripts/zaplog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
)

var once sync.Once

func main() {
	tool.InitLog()
	// name: 该服务的全局 trace
	// TODO: 整合 grafana 和 prometheus
	tracer.InitGlobal("http-gateway", zaplog.GetLogger())
	InitRPC()
	r := gin.New()
	r.Use(Tracer())

	r.Use(gin.Recovery())
	r.GET("/hello", Hello)
	r.Run(":7777")
}

// Hello api
func Hello(c *gin.Context) {
	var (
		name string
		ctx  context.Context
	)
	// TODO: 这里是为了将 ctx 传递进 client
	if pctx, ok := c.Get("ctx"); ok {
		ctx = pctx.(context.Context)
	} else {
		panic("c has not yet set ctx")
	}

	if name = c.Query("name"); name == "" {
		panic("name cannot be empty")
	}

	if span := opentracing.SpanFromContext(ctx); span != nil {
		newSpan := opentracing.StartSpan("SayHello", opentracing.ChildOf(span.Context()))
		defer newSpan.Finish()
		newSpan.SetTag("type", "SayHello client type")
		ctx = opentracing.ContextWithSpan(ctx, newSpan)
	}

	in := &pb.HelloRequest{Name: "mamengli"}
	resp, err := GetServerIns().SayHello(ctx, in)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"msg": resp.Message})
}

// InitRPC 服务发现
func InitRPC() {
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

// Tracer trace begin, root span
func Tracer() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// operationName: 标记一个 span, SpanKindRPCClient 标记是 client 侧
		span := opentracing.GlobalTracer().StartSpan(path, ext.SpanKindRPCClient)
		defer span.Finish()

		ext.HTTPUrl.Set(span, path)
		ext.HTTPMethod.Set(span, c.Request.Method)

		// TODO: 注意，这里设置的超时时长为 6s
		timeOut, cancel := context.WithTimeout(context.Background(), time.Second*6)
		defer cancel()

		// 将 span 注入 ctx 中, key: contextKey{}, value: span
		ctx := opentracing.ContextWithSpan(timeOut, span)
		c.Set("ctx", ctx)
		c.Next()

		// 响应
		ext.HTTPStatusCode.Set(span, uint16(c.Writer.Status()))
	}
}
