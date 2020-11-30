package demo

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
)

// [https://studygolang.com/articles/23461](Go: Finalizers)

/*
	SetFinalizer 用于用户优雅关闭后台 goroutine；
	1. 函数原型
	func SetFinalizer(x, f interface{})

	SetFinalizer 将对象 x 的终止器设置为 f，当垃圾收集器发现一个不能接触的（即引用计数为零，程序中不能再直接或间接访问该对象）具有终止器的对象时，
	它会清理该关联（对象-终止器）并在独立 go 程中调用 f(x)，这使 x 再次可以接触，但没有了绑定的终止器。
	如果 SetFinalizer 没有被再次调用，下一次垃圾收集器将视 x 为不可接触的，并释放 x。

	SetFinalizer(x, nil)会清理任何绑定到x的终止器。

	参数 x 必须是一个指向通过 new 申请的对象的指针，或者通过对复合字面值取址得到的指针。参数 f 必须是一个函数，它接受单个可以直接用 x 类型值赋值的参数，
	也可以有任意个被忽略的返回值。如果这两条任一条不被满足，本函数就会中断程序。

	2. 注意点
	- 即使程序正常结束或退出，但是对象被 gc 选中并回收之前，SetFinalizer 都不会执行，所以不要在 SetFinalizer 中执行 flush 这种操作；
	- SetFinalizer 最大的问题是延长了对象的生命周期，在第一次回收时执行 finalizer 函数，且目标对象重新变成可达状态，直到第二次才真正销毁，
	  这对于有大量对象分配的高并发计算场景，可能会造成很大麻烦；
	- 指针构成的 "循环引⽤" 加上 runtime.SetFinalizer 会导致内存泄露；
*/

type Foo struct {
	a int
}

func NewFoo(i int) *Foo {
	f := &Foo{a: rand.Intn(50)}
	runtime.SetFinalizer(f, func(f *Foo) {
		_ = fmt.Sprintf("foo " + strconv.Itoa(i) + " has been garbage collected")
	})
	return f
}

// 无保障性
// 在程序无法获取到一个 obj 所指向的对象后的任意时刻，finalizer 被调度运行，且无法保证 finalizer 运行在程序退出之前。
// 因此一般情况下，因此它们仅用于在长时间运行的程序上释放一些与对象关联的非内存资源。
func demo1() {
	for i := 0; i < 3; i++ {
		f := NewFoo(i)
		fmt.Println("xxx", f.a)
	}
	runtime.GC()
	time.Sleep(5 * time.Second)
}

func demo2() {
	// 垃圾收集比率：新分配内存和上次 GC 回收后剩下的实时数据比
	// 当新分配的数据与先前收集后剩余的实时数据达到此百分比时，会触发收集。
	// 默认值是GOGC = 100, 设置 GOGC = off 将完全禁用垃圾收集器, 必须手动出发才会 gc;
	debug.SetGCPercent(-1)

	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	// HeapObjects 分配的对象数
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1204), ms.HeapObjects)

	for i := 0; i < 1000000; i++ {
		f := NewFoo(i)
		_ = fmt.Sprintf("%d", f.a)
	}

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1204), ms.HeapObjects)

	runtime.GC()
	time.Sleep(time.Second)

	runtime.ReadMemStats(&ms)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1204), ms.HeapObjects)

	runtime.GC()
	runtime.GC()
	time.Sleep(5 * time.Second)
	fmt.Printf("Allocation: %f Mb, Number of allocation: %d\n", float32(ms.HeapAlloc)/float32(1024*1204), ms.HeapObjects)
}

// --------------------------- demo3
type Cache = *wrapper

type cache struct {
	contant   string
	stop      chan struct{}
	onStopped func()
}

func newCache() *cache {
	return &cache{
		contant: "test cache",
		stop:    make(chan struct{}),
	}
}

func NewCache() Cache {

	w := &wrapper{newCache()}
	go w.cache.run()

	runtime.SetFinalizer(w, (*wrapper).stop)
	return w
}

func (c *cache) run() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("clean...")
		case <-c.stop:
			if c.onStopped != nil {
				c.onStopped()
			}
			return
		}
	}
}

type wrapper struct {
	*cache
}

func (w *wrapper) stop() {
	close(w.cache.stop)
}

func demo3() {

}
