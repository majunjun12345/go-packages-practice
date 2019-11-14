package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	https://blog.csdn.net/preyta/article/details/80655736

	bufio 是缓冲区，达到条件后向底层数据读或写
	bytes.buffer 只是存放数据的一个空间, 创建：buf := new(bytes.Buffer)

	bufio.NewReaderSize 返回的 reader，可以 read readByte readBytes readLine readRune readString

	scan 可以实现按 行 字节 字符串 单词读，由于其他方法
*/

func main() {
	testRead()
	// testWrite()
	// testScan()
}

// read 是从 os.reader 读取数据
func testRead() {
	f, err := os.Open("Dockerfile")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	buf := bufio.NewReaderSize(f, 100)
	for {
		// data := make([]byte, 101)
		// n, err := buf.Read(data)

		data, err := buf.ReadString('\n')       // 相当于 readLine，注意这里是单引号
		line := strings.TrimSpace(string(data)) // 去掉 space
		if line != "" {                         // 有可能读到空行
			fmt.Println(line)
		}

		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
	}

	/*
		test writeto
	*/
	destF, err := os.Create("Dockerfile_bark")
	defer destF.Close()
	checkerr(err)

	src, err := os.Open("Dockerfile")
	defer src.Close()

	bufF := bufio.NewReaderSize(src, 100)
	n, err := bufF.WriteTo(destF)

	fmt.Println(n)
	checkerr(err)
}

// writer 是向 os.writer 写入数据
func testWrite() {
	// buf := bytes.NewBuffer(make([]byte, 0)) // 和 bytes.buffer 结合使用
	bufW := bufio.NewWriterSize(os.Stdout, 10)

	n, err := bufW.WriteString("mame") // 达到 10 字节后会自动向 os.stdout 中写入数据
	// n, err = bufW.WriteString("ma")
	bufW.Flush() // 将未达到指定字节数的数据写入 writer 中
	checkerr(err)
	fmt.Println(n)
}

// 有的人觉得使用 scan 要比 reader 好
func testScan() {
	f, err := os.Open("Dockerfile")
	defer f.Close()
	checkerr(err)

	// scan := bufio.NewScanner(strings.NewReader("abcdefghi"))
	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanLines) // ScanLines  ScanRunes  ScanBytes  ScanWords

	for scan.Scan() {
		fmt.Println(scan.Text())
	}
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
