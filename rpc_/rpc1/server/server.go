package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

/*
	&：取址符号，表示某个变量的地址，如：＆ａ
	＊：指针运算符，可以表示变量的指针**类型**，也可以表示一个指针变量所指向的存储单元，也就是这个地址对应的值(取指针对应的值)
*/

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B // 不甚明白，reply 是　*int，但是　*reply　是？
	return nil
}

func (t *Arith) Divide(args *Args, reply *Quotient) error {
	if args.B == 0 {
		return errors.New("divided by zero!")
	}
	reply.Quo = args.A / args.B
	reply.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	log.Fatal(http.ListenAndServe(":1234", nil))
}
