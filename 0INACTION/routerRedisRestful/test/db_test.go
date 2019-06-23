package test

import (
	"fmt"
	"testGoScripts/0INACTION/routerRedisRestful/db"
	"testGoScripts/0INACTION/routerRedisRestful/models"
	"testing"
)

/*
	测试文件以 _test 结尾,比如: db_test.py
	go test 是带缓存的，go test -count=1 可以去除缓存, 表示重复执行次数
	go test 的函数顺序执行

	单元测试:
		测试函数以 Test 开头,唯一参数: t *testing.T

		测试命令:
			go test -v xxx_test.go: 测试整个文件
			go test -v xxx_test.go -test.run TestInsert: 测试文件中的某个函数
			-v: 显示详细信息，可以打印内容

	基准测试:运行性能及耗费 CPU 的程度
		测试函数以 Benchmark 开头, 唯一参数: b *testing.B

		测试命令:
			go test -bench=. xxx_test.go: 表示执行该文件的全部压力测试函数, 加上 -v 就会执行单元测试
			-bench=.: 表示运行 xxx_test.go 文件中的所有基准测试
			-bench=BenchmarkInsert: 指定测试函数
			b.N=20000: 来设置压力次数
			-cpu=1,2,4,8: 控制 cpu 执行核数
			-benchmem: 显示内存分配情况
			-benchtime=5s: 控制执行时长
				基准测试的单次执行时间默认是 1s，如果函数执行时长未超过 1s，则会在 1s 内尽可能多的执行测试函数；

		关于计时：
			benchmark 函数开始，StartTimer 就开始计数，StopTimer 可以停止计数过程，
			再调用 StartTimer 可以从新开始计数。ResetTimer 可以重置计数器的次数；
			计数器内部不仅包含耗时数据，还包括内存分配的数据。
*/

func TestInsert(t *testing.T) {
	u := models.User{UserName: "menglima", Email: "15527254815@qq.com"}
	err := db.Insert(&u)
	if err != nil {
		fmt.Println(err)
		fmt.Println("end")
	}
}

func TestFindByID(t *testing.T) {
	i := 15
	u, err := db.FindOneByID(i)
	if err != nil {
		t.Error("err in find")
	}
	fmt.Println("u:", u)
	if u.UserName != "menglima" {
		t.Fatal("err")
	}
}

func TestDel(t *testing.T) {
	i := 1
	err := db.Delete(i)
	if err != nil {
		t.Fatal("err")
	}
}

func TestUpdate(t *testing.T) {
	u := models.User{Id: 1, UserName: "masanqi", Email: "111@qq"}
	i := 1
	user, err := db.Update(i, u)
	if err != nil {
		t.Fatal("err")
	}
	fmt.Println(user)
}

func BenchmarkInsert(b *testing.B) {
	b.StopTimer()
	b.StartTimer()

	b.N = 1000 // 指定次数

	for i := 0; i < b.N; i++ {
		u := models.User{UserName: "menglima", Email: "15527254815@qq.com"}
		err := db.Insert(&u)

		if err != nil {
			fmt.Println("err:", err)
			b.Error("operation fail!")
		}
	}
}

func BenchmarkFindById(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := db.FindOneByID(i)
		if err != nil {
			b.Error("err in find")
		}
	}
}
