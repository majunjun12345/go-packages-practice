实现了
* 服务注册
* 命名解析
* 服务发现
* 负载均衡
* 集成可视化实时运行时统计: statsviz 和 pprof
    [你喜欢的 Go 第三方库：一步为系统集成可视化实时运行时统计](https://polarisxu.studygolang.com/posts/go/pkg/statsviz/)
    [通过这个 Runtime 统计信息可视化库学到了什么？](https://mp.weixin.qq.com/s/sRxKJa-zKo-Lb7KRKLhthA)
* 使用 grpcui 进行 grpc 调用
* 整合 trace jeager: http://127.0.0.1:16686
  http-gateway(start trace: root span) => grpc client(span) => grpc server(span)
  在 gin => grpc client 时特别要注意 ctx 的传递, timeout 是在 r.Use(Tracer()) 传递进去的