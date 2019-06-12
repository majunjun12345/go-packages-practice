package main

import (
	"io"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go HandleConn(conn) // 加个 go 就能实现并发
	}
}

func HandleConn(con net.Conn) {
	defer con.Close() // 也要注意关闭
	for {
		_, err := io.WriteString(con, time.Now().String()+"\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second * 1)
	}
}
