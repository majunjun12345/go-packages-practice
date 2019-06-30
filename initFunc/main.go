package main

import "fmt"

/*
	init 函数:
		在包级别被定义
		仅执行一次的计算, 不管包被导入多少次, 都只会被初始化一次
		有几个 init 函数就会执行几个 init 函数
		初始化那些不能被初始化表达式完成初始化的变量
		不需要参数, 也没有返回值, 也无法被引用

	和初始化变量的执行顺序:(不建议依赖 init 的初始化顺序)
		当包中有谷歌文件,每个文件都有 初始化变量 和 init 函数时,
		1. 先初始化变量, 初始化变量的顺序是文件名的词法顺序
		2. 初始化变量完成之后开始初始化 init 函数, 也是根据文件名的词法顺序

	在 project 当中的使用:
		每个包如有需要运行前初始化的变量,可以在包级别定义 Init 函数进行变量的初始化
		在 app 包中定义 init 函数, 调用各个包的 Init 函数进行变量的一致性初始化
		(
			DB  定义全局变量, 通过 Init -> init 函数进行初始化
			collection 和 DB 一样(mongodb 中)
		)
*/

var _ int64 = s()

func init() {
	fmt.Println("init in sandbox.go")
}

func s() int64 {
	fmt.Println("calling s() in sandbox.go")
	return 1
}

func main() {
	fmt.Println("main")
}
