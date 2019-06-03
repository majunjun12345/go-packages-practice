package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	Id      string
	Balance uint64
}

/*
	NewEncoder 参数是 writer
	Encode 参数是结构体，将结构体转换为 json 写入至 writer 中
	Post 参数也是 reader

	new(bytes.Buffer) 是 readerwriter
*/

func main() {
	u := User{
		Id:      "1",
		Balance: 59,
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(u)
	fmt.Println(buf)
	res, _ := http.Post("https://httpbin.org/post", "application/json:charset=utf-8", buf)
	io.Copy(os.Stdout, res.Body)
}
