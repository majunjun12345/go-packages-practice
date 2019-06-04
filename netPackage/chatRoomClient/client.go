package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var sendMessage, recMessage = make([]byte, 1024), make([]byte, 1024)

func main() {
	reader := bufio.NewReader(os.Stdin)

	conn, err := net.Dial("tcp", "localhost:8080")
	defer conn.Close()
	CheckErr(err)

	_, err = conn.Write(sendMessage)
	CheckErr(err)

	fmt.Println("now begin to talk!")

	go read(conn)

	for {
		sendMessage, _, _ := reader.ReadLine()
		if string(sendMessage) == "quit" {
			fmt.Println("quit communication")
			os.Exit(1)
		}
		_, err := conn.Write(sendMessage)
		CheckErr(err)
	}
}

func read(conn net.Conn) {
	for {

		length, err := conn.Read(recMessage)
		CheckErr(err)
		fmt.Println(string(recMessage[:length]))
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
