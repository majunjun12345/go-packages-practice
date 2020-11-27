package zaplog

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

type Animal struct {
	Name string
	Age  int
}

func TestInitZapV2Logger(t *testing.T) {

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		t.Log("=======", err)
	// 	}
	// }()

	a := &Animal{
		Name: "cat",
		Age:  3,
	}

	b := make(map[string]interface{})
	b["key1"] = "value1"
	b["key2"] = "value2"

	NewLogger(
		SetInitialFields("=====", "xxxxxx"),
		SetLogFileDir("logs/"),
		SetDevelopment(true),
		SetLevel(zap.DebugLevel),
	)

	for i := 0; i < 1; i++ {
		// 结构体
		Debug("animal info", zap.Any("animal", a))
		// map
		Info("test map", zap.Reflect("some key and value", b))
		// string
		Warn(fmt.Sprint("warn log ", 3), zap.String("level", `{"a":"4","b":"5"}`))
		// int
		Error(fmt.Sprint("err log ", 4), zap.Int("int", 5))

		// ...

		// ctx := context.WithValue(context.Background(), tracer.LogTraceKey, "46b1506e7332f7c1:7f75737aa70629cc:3bb947500f42ad71:1")
		// lg.Info("=====", zap.Any("ctx", ctx))
	}

	// will recovery if set recovery, the testing is ok
	// Panic("panic")

	// will exit, the testing is failing； use this one，please Comments Panic first
	// Fatal("fatal")

	tracer, closer, err := NewJaegerTrace("hello-world-trace", "127.0.0.1:6831")
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	defer closer.Close()

	// 第一级 span
	span := tracer.StartSpan("say-hello-span")
	defer span.Finish()
	span.SetTag("hello-to", "helloTo")
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	// sp := opentracing.StartSpan("operation_name")
	// defer sp.Finish()
	// ctx := opentracing.ContextWithSpan(context.Background(), sp)

	Info("test map", GetTraceFields(ctx)...)
}

func NewJaegerTrace(serviceName string, jaegerHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		// 采集器
		Sampler: &config.SamplerConfig{
			Type:  "const", // const 固定采样
			Param: 1,       // 1=全采样，0=不采样
		},
		// 记录器
		Reporter: &config.ReporterConfig{
			// 配置jaeger Agent的ip与端口，以便将tracer的信息发布到agent中，6381 接口是接受压缩格式的thrift协议数据
			LocalAgentHostPort: "127.0.0.1:6831",
			LogSpans:           true,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		fmt.Printf("ERROR: cannot init Jaeger: %v\n", err)
	}

	opentracing.SetGlobalTracer(tracer) // 用于产生后面的子 span
	return tracer, closer, err
}
