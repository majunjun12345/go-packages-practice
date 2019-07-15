package main

import (
	"io"

	"bufio"
	"fmt"
	"os"
)

/*
	打开文件不会耗费资源,打开文件只是使用文件句柄,并不会将文件导入内存

	流是处理大文件网络传输的最好办法，参看 io.copy 和 io.Pipes
*/

func main() {

}

// 流式读取文本,因为有换行符
func StreamRead() {
	f, err := os.Open("filename")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		fmt.Println(scan.Text())
	}
}

// 分片处理,针对二进制文件
func SliceRead() {
	f, err := os.Open("filename")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	s := make([]byte, 4096)
	for {
		_, err := f.Read(s)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Fprintln(os.Stdout, s)
	}
}
