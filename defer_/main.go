package main

import (
	"fmt"
	"os"
)

/*
	defer return 返回值：
	主要的区别在于 返回参数的命名与否
	匿名返回参数，defer 内对 返回参数的值进行修改，不会影响返回值
	命名返回参数，defer 内对 返回参数的值进行修改，会影响返回值；defer 函数好想不能有返回值，但是貌似可以 return

	defer return 的执行顺序：
	defer 不能放在 return 后面，否则不会执行！
*/

func main() {
	// testOrderA() // 1 2 defer
	// testOrderB() // 1
	fmt.Println("main:", testOrderC()) // main: JSON: internal error: open test.test: no such file or directory

	// fmt.Println(testDeferA()) // 0
	// fmt.Println(testDeferB()) // 1
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
