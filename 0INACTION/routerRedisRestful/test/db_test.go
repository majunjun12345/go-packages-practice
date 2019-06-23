package test

import (
	"fmt"
	"testGoScript/0INACTION/routerRedisRestful/db"
	"testGoScript/0INACTION/routerRedisRestful/models"
	"testing"
)

/*
	测试文件以 _test 结尾,比如: db_test.py

	单元测试:
		测试函数以 Test 开头,唯一参数: t *testing.T

		测试命令:
			go test -v xxx_test.go: 测试整个文件
			go test -v xxx_test.go -test.run TestInsert: 测试文件中的某个函数

	基准测试:运行性能及耗费 CPU 的程度
		测试函数以 Benchmark 开头, 唯一参数: b *testing.B

		测试命令:
			go test -bench=. xxx_test.go: 表示执行该文件的全部压力测试函数, 加上 -v 就会执行单元测试
			-bench=.: 表示运行 xxx_test.go 文件中的所有基准测试
			-bench=BenchmarkInsert: 指定测试函数
			b.N=20000: 来设置压力次数
			-cpu=3: 控制 cpu 执行核数
			-benchtime=5s: 控制执行时长
			-benchmem: 显示内存分配情况
*/

func TestInsert(t *testing.T) {
	u := models.User{UserName: "menglima", Email: "15527254815@qq.com"}
	err := db.Insert(&u)
	if err != nil {
		fmt.Println(err)
		fmt.Println("end")
	}
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
