package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// https://www.cntofu.com/book/3/zh/03.4.md 终结贴

// https://baijiahao.baidu.com/s?id=1628393222631709276&wfr=spider&for=pc

/*

	recover:
	recover 一般放在 defer 函数中捕获错误，必须显式调用,只能捕获最后一个错误；

	defer 不能放在 return 后面，否则不会执行！

	defer return 的执行顺序：
	给返回值赋值
	执行 defer
	return
*/

func main() {
	// testOrderA() // 1 2 defer
	// testOrderB() // 1
	// fmt.Println("main:", testOrderC()) // main: JSON: internal error: open test.test: no such file or directory

	// fmt.Println(testDeferA()) // 0
	// fmt.Println(testDeferB()) // 1

	/*
		输出 222 但不输出 111, panic 后面的 defer 不会执行
	*/
	// deferPanic()

	// Tparams()
	// mergeFile()
	// f1()
	// f2()
	// f3()

	// fmt.Println(t1())
	// fmt.Println(t2())
	// fmt.Println(t3())
	// fmt.Println(t4())
	// fmt.Println(t5())

	t6()
}

func testOrderA() {
	fmt.Println("1")
	defer func() {
		fmt.Println("defer")
	}()
	fmt.Println("2")
	return
}

func testOrderB() {
	fmt.Println("1")
	return
	defer func() {
		fmt.Println("defer")
	}()
	fmt.Println("2")
}

func testOrderC() (err error) {
	fi, err := os.Open("test.test")

	defer func() {
		fi.Close()
		fmt.Println("close")

		if p := recover(); p != nil {
			fmt.Println("p")
			err = fmt.Errorf("JSON: internal error: %v", p)
		}
	}()

	if err != nil {
		panic(err)
	}
	return

}

// ------------------------------------------------------------

func testDeferA() int { // 匿名返回参数
	i := 0
	defer func() {
		i++
	}()
	return i
}

func testDeferB() (i int) { // 命名返回参数
	i = 0
	defer func() {
		i++
	}()
	return
}

func deferPanic() {
	defer func() {
		fmt.Println("222")
	}()
	panic("panic")
	defer func() {
		fmt.Println("111")
	}()
}

// ------------------------------------------------------------
/*
	每次defer语句执行的时候，会把函数“压栈”，函数参数会被拷贝下来;

	在defer函数定义时，对外部变量的引用是有两种方式的，分别是作为函数参数和作为闭包引用。
	作为函数参数，则在defer定义时就把值传递给defer，并被cache起来，是拷贝，不会随着以后的更改而改变；(感觉也是引用)，只有在执行裸代码块的时候才是赋值, 详情见下面两个函数
	作为闭包引用，则会在defer函数真正调用时根据整个上下文确定当前的值。

	defer语句并不会马上执行，而是会进入一个栈，函数return前，会按先进后出的顺序执行;

	return xxx 和 defer 的执行顺序：1 和 3 才是 return 真正执行的命令
	1. 返回值 = xxx
	2. 调用defer函数
	3. 空的return
*/

func Tparams() {
	t := 6

	// 1 和 2 一样，和 3 不一样
	// 1
	defer func() {
		fmt.Println(t) // 9  "引用" 必包
	}()
	t = t + 1
	// 1
	defer func(i int) {
		fmt.Println(t) // 9  "引用"  函数
	}(t)
	t = t + 1
	// 3
	defer fmt.Println(t) // 8  "传值"  函数
	t = t + 1
}

// 不明白上下两个函数的执行机制不一样，下面是传值，上面是传址
func mergeFile() error {
	fmt.Println("begin")
	f, _ := os.Open("file1.txt")
	if f != nil {
		defer func(f io.Closer) {
			fmt.Println("222", f) // 222 &{0xc00004c120}
			if err := f.Close(); err != nil {
				fmt.Printf("defer close file1.txt err %v\n", err)
			}
		}(f)
	}
	// ……
	f, _ = os.Open("file2.txt")
	if f != nil {
		defer func(f io.Closer) {
			fmt.Println("111", f) // 111 &{0xc000094000}
			if err := f.Close(); err != nil {
				fmt.Printf("defer close file2.txt err %v\n", err)
			}
		}(f)
	}
	return nil
}

func f1() {
	var err error
	defer fmt.Println(err) // nil 函数
	err = errors.New("defer error")
	return
}
func f2() {
	var err error
	defer func() { // defer error 必包
		fmt.Println(err)
	}()
	err = errors.New("defer error")
	return
}
func f3() {
	var err error
	defer func(err error) {
		fmt.Println(err) // 函数 nil
	}(err)
	err = errors.New("defer error")
	return
}

// --------------------------------------------读取外部变量  对返回值的改变
/*
	使用函数时, 不会受后续参数的改变的影响
	使用闭包时, 随着后续参数的改变而改变

	匿名返回, 返回值不受 defer 的影响
	命名返回: 使用函数, 不会改变返回值, 使用闭包, 改变返回值;

	总: 函数不会影响匿名命名返回值,函数内变量也不会受后续变量的改变影响; 闭包影响命名返回值, 不影响匿名返回值, 闭包内变量随后续的改变而改变;
*/

func t1() int { // 4
	n := 3
	defer func(i int) {
		fmt.Println("defer:", i) // 3  函数, 将上面的值拷贝
		i++
	}(n)
	n++
	return n
}

func t2() int { // 4
	n := 3
	defer func() {
		fmt.Println("defer:", n) // 4 闭包, 值相当于是 "引用", 受下面的影响
		n++
	}()
	n++
	return n
}

func t3() int { // 4
	n := 3
	defer fmt.Println("defer:", n) // 3 函数, 将上面的值拷贝
	n++
	return n
}

// -------------------------- 命名返回参数
func t4() (n int) { // 4, 函数传值后就变成另外的变量了
	n = 3
	defer func(i int) {
		fmt.Println("defer:", i) // 3 函数
		i++
	}(n)
	n++
	return
}

func t5() (n int) { // 5
	n = 3
	defer func() {
		fmt.Println("defer:", n) // 4
		n++
	}()
	n++
	return n
}

// ----------------------------------------------------------------
//闭包，defer 能够在函数开始执行前 获取宿主函数的变量执行一些初始化函数
func t6() {
	a := "mamengli"
	defer fmt.Println(hello(a)(" menglima"))
}

type fn func(string) string

func hello(s string) fn {
	fmt.Println(s)
	return func(z string) string {
		fmt.Println("end")
		return s + z
	}
}
