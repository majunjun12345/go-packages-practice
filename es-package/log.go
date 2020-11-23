package main

import "fmt"

//tracelog 实现 elastic.Logger 接口
type tracelog struct{}

//实现输出
func (tracelog) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
