package main

import (
	"fmt"
	"log"
	"net/rpc"
	"testGoScripts/rpc_/tcpRpc"
)

func main() {

	// 和 http 唯一的差别是 Dial 和 DialHTTP
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// synchronous call
	args := &tcpRpc.Args{17, 8}

	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot tcpRpc.Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

	args1 := &tcpRpc.Args{8, 5}
	var reply1 int
	err = client.Call("Arith.Minus", args1, &reply1)
	fmt.Printf("Arith: %d-%d=%d\n", args1.A, args1.B, reply1)
}
