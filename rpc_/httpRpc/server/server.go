package main

import (
	"log"
	"net/http"
	"net/rpc"
	"testGoScripts/rpc_/rpc1"
)

/*
	&：取址符号，表示某个变量的地址，如：＆ａ
	＊：指针运算符，可以表示变量的指针**类型**，也可以表示一个指针变量所指向的存储单元，也就是这个地址对应的值(取指针对应的值)
*/

func main() {
	arith := new(rpc1.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	log.Fatal(http.ListenAndServe(":1234", nil))
}
