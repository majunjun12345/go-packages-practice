package main

import (
	"fmt"
	"net"
)

var clients = make(map[string]net.Conn, 10)

func main() {
	data := make([]byte, 1024)

	listen, err := net.Listen("tcp", "localhost:8080")
	CheckErr(err)
	defer listen.Close()
	fmt.Println("listening at 8080...")

	for {
		conn, err := listen.Accept()
		CheckErr(err)
		clientIP := conn.RemoteAddr().String()
		clients[clientIP] = conn
		Notify(conn, " enter home!")

		go func(conn net.Conn) {
			length, err := conn.Read(data)
			if err != nil {
				fmt.Printf("client %s quit\n", clientIP)
				conn.Close()
			}
			fmt.Printf("%s enter room\n", clientIP)

			for {
				length, err = conn.Read(data)
				if err != nil { // EOF
					fmt.Printf("client %s quit\n", clientIP)
					conn.Close()
					delete(clients, clientIP)
					Notify(conn, "has left home!")
					return // 如果没有 return，将会一直循环
				}
				res := string(data[:length])
				sprdMsg := clientIP + " said:" + res
				fmt.Println(sprdMsg)

				ret := "you said:" + res
				conn.Write([]byte(ret))
				fmt.Println("server send:", ret)
			}
		}(conn)
	}
}

func Notify(conn net.Conn, note string) {
	for client := range clients {
		leaveCli := conn.RemoteAddr().String()
		if client != leaveCli {
			clients[client].Write([]byte(leaveCli + " " + note))
		}
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(nil)
	}
}
