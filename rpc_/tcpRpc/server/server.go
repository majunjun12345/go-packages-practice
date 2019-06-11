package main

import (
	"log"
	"net"
	"net/rpc"
	"testGoScripts/rpc_/tcpRpc"
)

func main() {
	arith := new(tcpRpc.Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		rpc.ServeConn(conn)
	}
}
