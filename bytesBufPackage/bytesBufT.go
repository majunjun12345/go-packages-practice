package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	bytes strings 是同类型的包

	只要读出，就不能回头，string() bytes() 除外

	bytes.Buffer 本身就是一个缓存（内存块或容器），底层数据是 []byte，缓存的容量会根据需要
	自动调整。大多数情况下，使用 new(bytes.Buffer) 就可以初始化一个 Buffer 了，具备读写功能；

	write 向尾部写
	read 从游标头部读

	作用：
	    可以代替 + 做字符串 拼接
*/

func main() {

	buf := bytes.NewBuffer(make([]byte, 0, 3)) // 可读可写
	fmt.Println("len 0", buf.Len())            //0
	_, err := buf.Write([]byte("hello"))
	if err != nil {
		panic(err)
	}
	fmt.Println("len 1", buf.Len()) // 5
	fmt.Println(buf.Bytes())        // [104 101 108 108 111]
	fmt.Println("len 2", buf.Len()) // 5

	for { // 一般是通过 for 循环读取，直到遇到 EOF 表示全部读完
		_, err := buf.ReadByte()

		if err != nil && err != io.EOF {
			panic(err)
		}

		if err == io.EOF {
			break
		}
	}
	fmt.Println("len 3", buf.Len())          // 0，上面 hello 被读完，不再计入 len
	fmt.Println("len buf", len(buf.Bytes())) // 0，和上面一样

	buf.Write([]byte("world"))
	fmt.Println("remain:", buf.String()) // world 未读完，剩下的，但是游标不会变 和 Bytes() 一样

	buf.Write([]byte("good"))
	buf.WriteTo(os.Stdout)                // worldgood 将剩下的读取到 io.writer 中，游标挪到最后
	fmt.Println("\nstring:", buf.Bytes()) // [] WriteTo 表示剩下的全部读完，游标移到最后
	fmt.Println(buf.Cap())                // 29 暂时还不清楚怎么扩容机制
	fmt.Println("len -1", buf.Len())      // 0

	buf.Write([]byte("morning"))
	fmt.Println(buf.Next(1))  // [109]
	fmt.Println(buf.Next(1))  // [111]
	fmt.Println(buf.String()) // rning 剩下的

	data, _ := buf.ReadBytes([]byte("i")[0])
	fmt.Println(data)         // [114 110 105] 包含 i
	fmt.Println(buf.String()) // ng 剩下的

	buf.ReadFrom(strings.NewReader("mamengli")) // 从 io.reader 对象中读取数据(文件)
	fmt.Println(buf.String())                   // ngmamengli

	buf.Truncate(3)           // 保留前 3 个
	fmt.Println(buf.String()) // ngm

	buf.Reset() // 清空 buf
}
