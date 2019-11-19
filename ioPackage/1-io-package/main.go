package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	io 相关的几个包：
		io
		io/ioutil
		bufio
		fmt

	实现了 reader 或 writer 的类型
		os.File(os.Stdout) 同时实现了 io.Reader 和 io.Writer
		strings.Reader 实现了 io.Reader
		bytes.Buffer 同时实现了 io.Reader 和 io.Writer
		bytes.Reader 实现了 io.Reader
		bufio.Reader/Writer 分别实现了 io.Reader 和 io.Writer


		compress/gzip.Reader/Writer 分别实现了 io.Reader 和 io.Writer
		crypto/cipher.StreamReader/StreamWriter 分别实现了 io.Reader 和 io.Writer
		crypto/tls.Conn 同时实现了 io.Reader 和 io.Writer
		encoding/csv.Reader/Writer 分别实现了 io.Reader 和 io.Writer
		mime/multipart.Part 实现了 io.Reader
		net/conn 分别实现了 io.Reader 和 io.Writer(Conn接口定义了Read/Write)
*/

/*
	reader:
	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	将 len(p) 个字节读取到 p 中, 返回读取的字节数 n
*/

func main() {
	// testAt()

	// FromTo()

	// seek()

	// copy()

	// multiR()

	MultiW()
}

//----------------- readerat writerat，在某处 读 写  offset
/*
	readxxx writexxx 都会从上一次结束位置继续
*/
func testAt() {
	r := strings.NewReader("my name is  mamengli")
	p := make([]byte, 6)
	r.ReadAt(p, 3)
	fmt.Println(string(p)) // name i； ReadAt 不会记住偏移量

	r.Read(p)
	fmt.Println(string(p)) // my nam；Read 会记住偏移量，下次从这之后读起
	r.Read(p)
	fmt.Println(string(p)) // e is  ；

	p2 := make([]byte, 100)
	n, err := r.Read(p2)
	fmt.Println(n, err) // 8 <nil>, err 也有可能是 EOF

	// -------------- writeAt
	fi, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fi.WriteString("hello")
	fi.WriteString("11111") // 从上一个结束的位置继续写
	fi.WriteAt([]byte(" world"), 0)

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("hahaha")
	writer.Flush() // 清空缓存区
}

// ReaderFrom 和 WriterTo, 实现一次性的 读 写
func FromTo() {
	fi, err := os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(fi) // 这里的 err 是 nil，非 io.EOF, 等同于 ioutil.ReadFilr()
	writer.Flush()

	// 实现了 reader 接口的对象都有 WriteTo 方法
	reader := strings.NewReader("hahaha")
	reader.WriteTo(os.Stdout)

	buf := bytes.NewBuffer([]byte("menglima"))
	buf.WriteTo(os.Stdout)

	r := bufio.NewReader(strings.NewReader("hahaha"))
	r.WriteTo(os.Stdout)
}

// seeker 设置文件读取 写入的偏移量
func seek() {
	r := strings.NewReader("12345")

	r.Seek(1, io.SeekStart)
	data, _ := r.ReadByte()
	fmt.Println(string(data))

	r.Seek(-2, io.SeekEnd)
	data, _ = r.ReadByte()
	fmt.Println(string(data))

	r.Seek(2, io.SeekCurrent)
	data, _ = r.ReadByte()
	fmt.Println(string(data))
}

// copy
func copy() {
	reader := strings.NewReader("123")
	writer := bufio.NewWriter(os.Stdout)
	io.Copy(writer, reader)
	writer.Flush() // 这里也必须要 flush

	// 下面这两个是一样的效果
	writer.ReadFrom(reader)
	reader.WriteTo(writer)
}

// multipart

func multiR() {
	readers := []io.Reader{
		strings.NewReader("strings reader"),
		bytes.NewBufferString("bytes reader"),
	}
	reader := io.MultiReader(readers...)

	for {
		buf := make([]byte, 10)
		_, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(string(buf))
	}
}

func MultiW() {
	fi, err := os.Create("tmp.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	writers := io.MultiWriter(fi, os.Stdout)
	writers.Write([]byte("hello"))
}
