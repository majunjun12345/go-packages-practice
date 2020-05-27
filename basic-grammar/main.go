package main

import (
	"fmt"
	"strings"
)

type Handler func(string) string

func process1(h Handler, s string) string {
	return "process1" + h(s)
}

func process2(h Handler, s string) string {
	return "process2" + h(s)
}

/*
	函数作为参数，且里面有 return
*/
func main() {
	processResult := process1(func(s string) string {
		return strings.ToUpper(s)
	}, "gch")
	fmt.Println(processResult)
}
