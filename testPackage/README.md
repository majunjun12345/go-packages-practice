
测试文件以 _test 结尾,比如: db_test.py
go test 是带缓存的，go test -count=1 可以去除缓存, 表示重复执行次数

单元测试:
    测试函数以 Test 开头,唯一参数: t *testing.T

    测试命令:
        go test -v xxx_test.go: 测试整个文件
        go test -v -test.run TestBubble: 测试文件中的某个函数
        -v: 显示详细信息，可以打印内容

表组测试：
    见详细例子

基准测试:运行性能及耗费 CPU 的程度
    测试函数以 Benchmark 开头, 唯一参数: b *testing.B

    测试命令:
        go test -bench=. : 表示执行该文件的全部压力测试函数, 加上 -v 就会执行单元测试
        go test -bench=. -test.run BenchmarkBubble：测试某一个函数
        -count=5：执行次数
        -bench=.: 表示运行 xxx_test.go 文件中的所有基准测试
        -bench=BenchmarkInsert: 指定测试函数
        b.N=20000: 来设置压力次数
        -cpu=1,2,4,8: 控制 cpu 执行核数
        -benchmem: 显示内存分配情况
            ns/op-nanosecond/operation 即执行一次操作消耗的时间。52.0 ns/op即平均每执行一次操作消耗0.052毫秒
            B/op 平均每次操作需要占用的内存空间（字节）
            allocs/op 平均每次操作需要分配内次的次数
        -benchtime=5s: 控制执行时长
            基准测试的单次执行时间默认是 1s，如果函数执行时长未超过 1s，则会在 1s 内尽可能多的执行测试函数；

    关于计时：
        benchmark 函数开始，StartTimer 就开始计数，StopTimer 可以停止计数过程，
        再调用 StartTimer 可以从新开始计数。ResetTimer 可以重置计数器的次数；
        计数器内部不仅包含耗时数据，还包括内存分配的数据。
    
    基准测试中,被测试函数必须是稳态的, 因为 b.N 会根据被测函数进行不断调整,直到达到稳定, 不过也可以强制设定 b.N 的值;

- t.SkipNow()
    必须写在函数首行, 用于暂时不用测试的函数

- sub test(子 test)
    测试文件内的函数并不保证顺序执行,如果测试文件内函数相互依赖,一定要保证顺序执行,怎么办呢?
    test case 开头字母大写才会被执行, 如果要使用 sub test,可以将开头字母改为小写

- 如果测试之前需要做初始化的东西呢?
    定义 TestMain 函数, 作为整个 test case 的入口, 做一些初始化工作,比如数据库连接文件打开登录等;
    里面必须有 M.Run(), 不然其他 test case 不会执行