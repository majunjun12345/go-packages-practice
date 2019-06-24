package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// os.Create("test.log")

	// os.Mkdir
	file, err := os.Open("Dockerfile")
	defer file.Close()
	CheckErr(err)

	content := make([]byte, 3)
	ret, err := file.Seek(10, 0)
	CheckErr(err)
	fmt.Println("ret", ret)

	file.Read(content) // 11 - 14
	fmt.Println("content:", string(content))

	io.Copy(os.Stdout, file) // 从 read 后开始，read 改变了 seek； 14 - ...

	file.Sync() // 同 flush，就是创建一个 bufio，先将文件写入 bufio，然后再将 bufio 中的内容写入到文件

	if _, err = os.Stat("docker"); os.IsNotExist(err) { // state fileinfo，IsNotExist 判断错误类型
		fmt.Println("file does not exit")
	}

	buf := bufio.NewWriter(os.Stdout)
	n, err := buf.WriteString("========")
	buf.Flush()
	fmt.Println(n, err)

	os.Exit(1)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// os.IsPermission 判断文件读写权限
// os.IsExist 判断问价是否存在
