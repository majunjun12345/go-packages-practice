package main

import (
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"
)

type HttpError struct {
	Code    int
	Desc    string
	Message string
}

func NewHttpError(code int, desc, message string) *HttpError {
	return &HttpError{
		Code:    code,
		Desc:    desc,
		Message: message,
	}
}

func (he *HttpError) Error() string {
	return fmt.Sprintf("errcode:%v, errdesc:%s, detail message:%s\n", he.Code, he.Desc, he.Message)
}

func main() {
	// errTest()

	// panicTest()

	E()
}

// err
func errTest() {
	newErr := NewHttpError(500, "internal server error", "out of index")
	fmt.Print(newErr)

	err := fmt.Errorf("something wrong has happend:%s", "out of index")
	fmt.Println(err)

	er := errors.New("internal server error")
	fmt.Println(er)

	// 可以打印出错误发生的文件和行号
	if er != nil {
		if _, file, line, ok := runtime.Caller(0); ok { // 0 表示调用层级，0，1 显示的是自己的，2 和 3 为包级别
			fmt.Println(file, line, ok) // /Users/majun/go/src/testGoScripts/errorPanicRecovery/main.go 46 true
		}
	}
}

// defer panic recovery
/*
	recovery:
		当发生 panic 后,使用 recovery 会捕获该异常,交给上层调用者,并正常执行剩下的代码;
	没有 recovery:
		如果没有 recovery, 函数遇到 panic 后会终止运行, 在执行完所有的延迟函数后, 程序控制返回到该函数的调用方,
		这样的过程会一直持续下去, 直到当前的协程的所有函数都返回退出,然后程序会打印出 panic 信息, 接着打印出堆栈跟踪,最后程序终止!
*/
func panicTest() {
	// defer panic("d") // recovery 执行后,继续panic,捕获不到了
	defer func() {
		if r := recover(); r != nil { // panic a 和 b, 只能捕获到 a(recovery 前的最后一个错误);
			fmt.Println("recovery from:", r)
			debug.PrintStack() // 打印出出调用栈
		}
	}()
	defer panic("a")
	// defer panic("b")
	fmt.Println("c")
}

// go1.13
func E() {
	err1 := errors.New("err one")
	err2 := fmt.Errorf("err2:[%w]", err1)
	err3 := fmt.Errorf("err3:[%w]", err2)
	fmt.Println(err3)
	fmt.Println(errors.Is(err3, err1))
	fmt.Println(errors.Unwrap(err3))
}
