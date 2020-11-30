[万字长文深入浅出 Golang Runtime](https://zhuanlan.zhihu.com/p/95056679)


- Gosched
让当前线程让出 cpu 以让其它线程运行,它不会挂起当前线程，因此当前线程未来会继续执行
- NumCPU
返回当前系统的 CPU 核数量
- GOMAXPROCS
设置最大的可同时使用的 CPU 核数
- Goexit
退出当前 goroutine(但是defer语句会照常执行)
- NumGoroutine
返回正在执行和排队的任务总数
- GOOS
目标操作系统

- Caller
获取调用信息