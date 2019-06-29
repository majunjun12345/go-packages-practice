package main

import (
	"fmt"
)

/*
	匿名函数
*/

func printMessage(message string) {
	fmt.Println(message)
}

func getPrintMessage() func(string) { // 返回函数需要写参数
	return printMessage
}

func testAnonymousFunc() {
	printMessage("menlgima")

	func(message string) {
		fmt.Println(message)
	}("mamengli")

	printFunc := getPrintMessage()
	printFunc("masanqi")
}

func main() {

	// testAnonymousFunc()

	// closure1("menglima")

	// closureFunc := closure2("mamengli")
	// closureFunc()

	// testClosure3()

	// testClousre4()

	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))

}

/*
	闭包: 定义在函数内部的函数, 能够访问外部变量

	闭包在外面再次调用时, 会记住外部变量的状态
*/

func closure1(name string) {
	test := "hello " + name

	foo := func() {
		fmt.Println(test)
	}

	foo()
}

func closure2(name string) func() {
	test := "hello " + name
	return func() {
		fmt.Println(test)
	}
}

// closure and state
func closure3(i int) (func() int, func() int) {
	one := func() int {
		return i
	}

	two := func() int {
		i++
		return i
	}

	return one, two
}

func testClosure3() {
	a, b := closure3(1)
	a2, b2 := closure3(1)
	fmt.Println(a(), b())   // 1, 2
	fmt.Println(a2(), b2()) // 1, 2
	fmt.Println(a(), b())   // 2, 3  再次调用返回的函数, 会记住上次变量的状态
}

// 陷阱
func closure4() []func() {
	a := []int{1, 2, 3, 4}

	var funcs []func()

	for i := range a {
		// for i := 0; i < len(a); i++ {   // 不能用这种方式, 因为闭包会记录变量的状态
		funcs = append(funcs, func() {
			fmt.Println(i, a[i])
		})
	}
	return funcs
}

func testClousre4() {
	res := closure4()
	for i := 0; i < len(res); i++ {
		res[i]() // 这里全部输出 3 4 也是同一个原因
	}
}

func fib(n int) (b int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return
}
