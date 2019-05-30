package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*
	fmt.Fprintf 向 writer 里面写数据

	可以在 conn 里面读取数据，也可以往里面写入数据
*/

func main() {
	conn, err := net.Dial("tcp", ":8080")
	CheckErr(err)
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		CheckErr(err)
		fmt.Fprintf(conn, text+"\n")
		message, err := bufio.NewReader(conn).ReadString('\n')
		CheckErr(err)
		fmt.Println("receive message from client:", message)
	}
}

func main1() {

	// connect to this socket
	conn, _ := net.Dial("tcp", ":8081")
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
