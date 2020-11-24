package tracer

import (
	"context"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/rpcmetrics"
	"github.com/uber/jaeger-lib/metrics"
	"github.com/uber/jaeger-lib/metrics/expvar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// InitGlobal create global trace
// JAEGER_AGENT_HOST=192.168.0.12
// JAEGER_AGENT_PORT=6831
// JAEGER_REPORTER_LOG_SPANS=1
// JAEGER_SAMPLER_TYPE=const
// JAEGER_SAMPLER_PARAM=1
func InitGlobal(name string, logger jaeger.Logger) {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(err)
	}
	cfg.ServiceName = name

	// TODO(ys) a quick hack to ensure random generators get different seeds, which are based on current time.
	time.Sleep(200 * time.Millisecond)

	metricsFactory := expvar.NewFactory(10).Namespace(metrics.NSOptions{Name: name, Tags: nil})
	// 集中式日志系统（Logging），集中式度量系统（Metrics）和分布式追踪系统（Tracing）
	tracer, _, err := cfg.NewTracer(
		config.Logger(logger),

		// TODO: 不清楚这两个是什么意思
		config.Metrics(metricsFactory),
		config.Observer(rpcmetrics.NewObserver(metricsFactory, rpcmetrics.DefaultNameNormalizer)),
	)
	if err != nil {
		panic(err)
	}
	opentracing.InitGlobalTracer(tracer)
}

var TracingComponentTag = opentracing.Tag{Key: string(ext.Component), Value: "ctx"}

//MDReaderWriter metadata Reader and Writer
type MDReaderWriter struct {
	metadata.MD
}

//ForeachKey range all keys to call handler
func (c MDReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, vs := range c.MD {
		for _, v := range vs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set implements Set() of opentracing.TextMapWriter
func (c MDReaderWriter) Set(key, val string) {
	key = strings.ToLower(key)
	c.MD[key] = append(c.MD[key], val)
}

//OpenTracingClientInterceptor  rewrite client's interceptor with open tracing
func OpenTracingClientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, resp interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		var parentCtx opentracing.SpanContext
		if parent := opentracing.SpanFromContext(ctx); parent != nil {
			parentCtx = parent.Context()
		}
		cliSpan := tracer.StartSpan(
			method,
			opentracing.ChildOf(parentCtx),
			TracingComponentTag,
			ext.SpanKindRPCClient,
		)
		defer cliSpan.Finish()
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}
		mdWriter := MDReaderWriter{md}
		// 注入 spanContext
		err := tracer.Inject(cliSpan.Context(), opentracing.TextMap, mdWriter)
		if err != nil {
			panic(err)
		}
		ctx = metadata.NewOutgoingContext(ctx, md)
		err = invoker(ctx, method, req, resp, cc, opts...)
		if err != nil {
			panic(err)
		}
		return err
	}
}

//OpentracingServerInterceptor rewrite server's interceptor with open tracing
func OpentracingServerInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		// 从 http 请求里面解析出上一个服务的span信息
		spanContext, err := tracer.Extract(opentracing.TextMap, MDReaderWriter{md})
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			panic(err)
		}
		serverSpan := tracer.StartSpan(
			info.FullMethod,
			ext.RPCServerOption(spanContext),
			TracingComponentTag,
			ext.SpanKindRPCServer,
		)
		defer serverSpan.Finish()
		ctx = opentracing.ContextWithSpan(ctx, serverSpan)
		return handler(ctx, req)
	}
}
