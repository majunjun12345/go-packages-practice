[CNCF 刚宣布 etcd 正式毕业：赶紧实战下 etcd 的服务发现](https://mp.weixin.qq.com/s/44duJY9BSS3Ki07jt4pgOA)

实现了
* 服务注册
* 命名解析
* 服务发现
* 负载均衡
* 集成可视化实时运行时统计: statsviz 和 pprof
    [你喜欢的 Go 第三方库：一步为系统集成可视化实时运行时统计](https://polarisxu.studygolang.com/posts/go/pkg/statsviz/)
    [通过这个 Runtime 统计信息可视化库学到了什么？](https://mp.weixin.qq.com/s/sRxKJa-zKo-Lb7KRKLhthA)
* 使用 grpcui 进行 grpc 调用
  [[系列] Go gRPC 调试工具](https://www.cnblogs.com/xinliangcoder/p/11825748.html)
  grpcui -plaintext 127.0.0.1:1234(grpc ip:port)
* 使用 grpcurl
  grpcurl --plaintext 127.0.0.1:1234 list
  grpcurl --plaintext 127.0.0.1:1234 list proto.Hello
  grpcurl --plaintext 127.0.0.1:1234 describe proto.Hello
  调用方法: `grpcurl -plaintext -d '{"name": "mamenglihahaha"}' localhost:1234 proto.Hello/SayHello`

* 整合 trace jeager: http://127.0.0.1:16686
  // JAEGER_AGENT_HOST=127.0.0.1
  // JAEGER_AGENT_PORT=6831
  // JAEGER_REPORTER_LOG_SPANS=1
  // JAEGER_SAMPLER_TYPE=const
  // JAEGER_SAMPLER_PARAM=1
  http-gateway(start trace: root span) => grpc client(span) => grpc server(span)
  在 gin => grpc client 时特别要注意 ctx 的传递, timeout 是在 r.Use(Tracer()) 传递进去的
  全链路：agent query es [Jaeger OpenTracing 体验](https://blog.csdn.net/nia305/article/details/90757235)

* 超时保活
  通过GRPC的重连机制可以完美的解决连接池设计原则中的空闲连接的超时与保活问题
  基于 GRPC 的 Keepalive 特性，我们不需要自己实现保活机制，也无需关注连接池中的连接是否有效，因为就算失效，GRPC会自动重连的；
  此时只不过耗时会略微增加，即认为除了服务器一直处于 down 状态等原因，连接池中的连接是始终有效的。