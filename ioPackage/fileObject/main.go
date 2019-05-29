package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.Open("test.txt")
	CheckErr(err)

	content := make([]byte, 10)
	n, err := fileObj.Read(content)
	CheckErr(err)
	fmt.Println(n, string(content))

	n, err = fileObj.ReadAt(content, 10)
	fmt.Println(string(content))

	/*
		只有 open 一个目录时,才会有 Readdir 方法
	*/
	fileObj2, err := os.Open("../ioutilPackage")
	fileinfos, err := fileObj2.Readdir(2)
	CheckErr(err)
	for _, findo := range fileinfos {
		fmt.Println(findo.Name())
	}

	/*
		0 表示从文件头开始偏移,1 表示从当前位置开始偏移, 2 表示从文件末尾开始偏移
		读取的是偏移量 15 后的数据
	*/
	n1, err := fileObj.Seek(15, 0)
	CheckErr(err)
	fmt.Println(n1)
	fileObj.Read(content)
	fmt.Println(string(content))

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
