package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// https://baijiahao.baidu.com/s?id=1628393222631709276&wfr=spider&for=pc

/*
	明白两点：
	defer 后面跟 函数 和 必包对变量的影响不一样
	return 后面带不带变量，defer 的影响也不一样
*/

/*
	defer return 返回值：
	主要的区别在于 返回参数的命名与否
	匿名返回参数，defer 内对 返回参数的值进行修改，不会影响返回值
	命名返回参数，defer 内对 返回参数的值进行修改，会影响返回值；defer 函数好想不能有返回值，但是貌似可以 return

	recover:
	recover 一般放在 defer 函数中捕获错误，必须显式调用；

	defer return 的执行顺序：
	defer 不能放在 return 后面，否则不会执行！
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

	Tparams()
	// mergeFile()
	// f1()
	// f2()
	// f3()
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
	t := []int{1}

	// 1 和 2 一样，和 3 不一样
	// 1
	defer func() {
		fmt.Println(t) // 7  "引用" 必包
	}()
	t = append(t, 2)
	// 1
	defer func(i []int) {
		fmt.Println(t) // 7  "引用"  函数
	}(t)
	t = append(t, 3)
	// 3
	defer fmt.Println(t) // 6  "传值"  函数
	t = append(t, 4)
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
