package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	rpc1 "testGoScripts/rpc_/httpRpc"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "server")
	}

	serverAddress := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// synchronous call
	args := &rpc1.Args{17, 8}

	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot rpc1.Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

	args1 := &rpc1.Args{8, 5}
	var reply1 int
	err = client.Call("Arith.Minus", args1, &reply1)
	fmt.Printf("Arith: %d-%d=%d\n", args1.A, args1.B, reply1)
}
