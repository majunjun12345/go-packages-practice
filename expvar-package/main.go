package main

import (
	"expvar"
	"io"
	"net/http"
	"strconv"
	"testGoScripts/expvar-package/monitor"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paulbellamy/ratecounter"
)

/*
runtime 包拥有各种功能，包括goroutine数量，设置逻辑线程数量，当前go版本，当前系统类型等等；
expvar 包可以监控服务运行各项指标和状态

expvar包为监控变量提供了一个标准化的接口，它以 JSON 格式通过 /debug/vars 接口以 HTTP 的方式公开这些监控变量以及我自定义的变量。
通过它，再加上metricBeat，ES和Kibana，可以很轻松的对服务进行监控。这里用 gin 把接口暴露出来，用别的web框架也都可以。

expvar 返回给了我我之前自定义的数据，以及它本身要默认返回的数据，比如memstats：
1、Alloc uint64 //golang语言框架堆空间分配的字节数
2、TotalAlloc uint64 //从服务开始运行至今分配器为分配的堆空间总 和，只有增加，释放的时候不减少
3、Sys uint64 //服务现在系统使用的内存
4、Lookups uint64 //被runtime监视的指针数
5、Mallocs uint64 //服务malloc的次数
6、Frees uint64 //服务回收的heap objects的字节数
7、HeapAlloc uint64 //服务分配的堆内存字节数
8、HeapSys uint64 //系统分配的作为运行栈的内存
9、HeapIdle uint64 //申请但是未分配的堆内存或者回收了的堆内存（空闲）字节数
10、HeapInuse uint64 //正在使用的堆内存字节数
10、HeapReleased uint64 //返回给OS的堆内存，类似C/C++中的free。
11、HeapObjects uint64 //堆内存块申请的量
12、StackInuse uint64 //正在使用的栈字节数
13、StackSys uint64 //系统分配的作为运行栈的内存
14、MSpanInuse uint64 //用于测试用的结构体使用的字节数
15、MSpanSys uint64 //系统为测试用的结构体分配的字节数
16、MCacheInuse uint64 //mcache结构体申请的字节数(不会被视为垃圾回收)
17、MCacheSys uint64 //操作系统申请的堆空间用于mcache的字节数
18、BuckHashSys uint64 //用于剖析桶散列表的堆空间
19、GCSys uint64 //垃圾回收标记元信息使用的内存
20、OtherSys uint64 //golang系统架构占用的额外空间
21、NextGC uint64 //垃圾回收器检视的内存大小
22、LastGC uint64 // 垃圾回收器最后一次执行时间。
23、PauseTotalNs uint64 // 垃圾回收或者其他信息收集导致服务暂停的次数。
24、PauseNs [256]uint64 //一个循环队列，记录最近垃圾回收系统中断的时间
25、PauseEnd [256]uint64 //一个循环队列，记录最近垃圾回收系统中断的时间开始点。
26、NumForcedGC uint32 //服务调用runtime.GC()强制使用垃圾回收的次数。
27、GCCPUFraction float64 //垃圾回收占用服务CPU工作的时间总和。如果有100个goroutine，垃圾回收的时间为1S,那么就占用了100S。
28、BySize //内存分配器使用情况


https://blog.csdn.net/jeffrey11223/article/details/78886923/

expvar.Get("foo-prefix").String()    //获取到的值是`"foo"`   不是`foo`!!!
最好还是使用expvar.NewXxx() 的原对象的Value()方式来取值就不会有问题了。
*/

var test = expvar.NewMap("Test")

func init() {
	test.Add("key1", 3)
	test.Add("key2", 4)
}

func main() {
	router := gin.Default()

	//接口路由，如果url不是/debug/vars，则用metricBeat去获取会出问题
	router.GET("/debug/vars", monitor.GetCurrentRunningStats)
	router.GET("/increment", increment)
	s := &http.Server{
		Addr:           ":8800",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

var (
	// 线程安全的频率统计
	counter       = ratecounter.NewRateCounter(1 * time.Minute)
	hitsperminute = expvar.NewInt("hits_per_minute")
)

func increment(c *gin.Context) {
	counter.Incr(1)
	hitsperminute.Set(counter.Rate())
	io.WriteString(c.Writer, strconv.FormatInt(counter.Rate(), 10))
}
