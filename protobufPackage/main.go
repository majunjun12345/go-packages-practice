package main

import (
	"fmt"
	"io"
	"os"
	gopb "testGoScripts/protobufPackage/proto"

	protobuf "github.com/golang/protobuf/proto"
)

func main() {
	// 使用 protobuf 写文件
	// WriteFile()

	// 使用 protobuf 读文件
	ReadFile()
}

func WriteFile() {
	//初始化protobuf数据格式
	msg := &gopb.HelloWorld{
		Id:   protobuf.Int32(17),
		Name: protobuf.String("BGbiao"),
		Opt:  protobuf.Int32(18),
	}

	filename := "./protobuf-test.txt"
	fObj, _ := os.Create(filename)
	defer fObj.Close()
	buffer, _ := protobuf.Marshal(msg)
	fObj.Write(buffer)
}

// 使用 protobuf 读文件
func ReadFile() {
	fileName := "protobuf-test.txt"
	fi, err := os.Open(fileName)
	checkError(err)
	fileInfo, err := fi.Stat()
	checkError(err)

	buf := make([]byte, fileInfo.Size())
	_, err = io.ReadFull(fi, buf)
	checkError(err)

	meg := &gopb.HelloWorld{}
	err = protobuf.Unmarshal(buf, meg)
	checkError(err)
	fmt.Printf("%+v\n", meg)

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
