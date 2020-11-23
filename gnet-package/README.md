[Gnet 导视](https://gnet.host/blog/presenting-gnet-cn/)

##### 目前支持以下六个事件

// PreWrite  预先写数据方法，在 server 端写数据回 client 端之前调用
func (s *echoServer) PreWrite() {
}

// Tick 服务器启动的时候会调用一次，之后就以给定的时间间隔定时调用一次，是一个定时器方法，设定返回的 delay 即可
func (s *echoServer) Tick() (delay time.Duration, action gnet.Action) {
	return
}

// React 当 server 端接收到从 client 端发送来的数据的时候调用。（你的核心业务代码一般是写在这个方法里）
func (s *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Println("===", c.RemoteAddr().String())

	s.pool.Submit(func() {
		time.Sleep(1 * time.Second)
		c.AsyncWrite(frame)
	})
	out = frame
	return
}

// OnInitComplete 当 server 初始化完成之后调用
func (s *echoServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Echo server is listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}

// OnShutdown 当服务被关闭的之后调用
func (s *echoServer) OnShutdown(svr gnet.Server) {
}

// OnOpened 当连接被打开的时候调用
func (s *echoServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

// OnClosed 当连接被关闭的之后调用
func (s *echoServer) OnClosed(c Conn, err error) (action Action) {
	return
}

type echoServer struct {
	*gnet.EventServer
	pool *goroutine.Pool
}

func main() {
	echo := &echoServer{
		pool: goroutine.Default(),
	}
	defer echo.pool.Release()
	log.Fatal(gnet.Serve(echo, "tcp://:9001", gnet.WithMulticore(true)))
}

通过 gnet.WithXXX 可以增加服务选项：
WithMulticore：使用多核来进行服务，核心数为 CPU 数量；
WithLoadBalancing：负载均衡策略
WithNumEventLoop：设置 event loop 数量
WithReusePort：使用端口复用特性，允许多个 sockets 监听同一个端口，然后内核会帮你做好负载均衡，每次只唤醒一个 socket 来处理 connect 请求，避免惊群效应；
WithTCPKeepAlive：
WithTicker：使用定时器
WithCodec：编解码器，支持自定义
WithLogger：

- Event Loop
  Event Loop是一个程序结构，用于等待和发送消息和事件；
  简单说，就是在程序中设置两个线程：一个负责程序本身的运行，称为"主线程"；
  另一个负责主线程与其他进程（主要是各种I/O操作）的通信，被称为"Event Loop线程"（可以译为"消息线程"）。
  就是程序本身于 IO 之间的桥梁；
  ![](http://www.ruanyifeng.com/blogimg/asset/201310/2013102004.png)