package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

/*
	可以在 conn 里面读取数据，也可以往里面写入数据

	一般的连接是阻塞至 Accept
	这里将 Accept 拿出来，通过 conn 不断获取数据
*/

func main() {
	listen, err := net.Listen("tcp", ":8080")
	CheckErr(err)

	conn, err := listen.Accept()
	CheckErr(err)

	for {
		content, isPrefix, err := bufio.NewReader(conn).ReadLine()
		CheckErr(err)
		fmt.Println(isPrefix) // false
		fmt.Printf("receive message from client:%s\n", string(content))
		newMessage := strings.ToUpper(string(content))
		conn.Write([]byte(newMessage + "\n")) // 后面的 "\n" 必须加上
	}

}

func main1() {

	fmt.Println("Launching server...")
	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")
	// accept connection on port
	conn, _ := ln.Accept()
	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
