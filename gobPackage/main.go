package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

/*
	为了让某个数据结构能够在网络上传输或能够保存至文件，它必须被编码然后再解码.
	现有编码格式：json xml google 的 protocol buffers

	gob 是 golang 自带的数据结构化编码/解码工具，最典型的使用场景就是 RPC，只能在 golang 中使用
	gob 和 json 类似，在发送端使用 encode 对数据进行编码，在接收端使用 decode 将数据解码为本地数据结构

	golang 可以利用 json 和 gob 来序列化 struct 对象，但 gob 编码能够实现 json 所不能支持的 struct 的方法的序列化。
	利用 gob 包序列化 struct 保存到本地十分简单。
*/

// 字段必须是可导出的
type P struct {
	X, Y, Z int
	Name    string
	B       bool
}

type Q struct {
	X, Y int
	Name string
	B    bool
}

func main() {
	// 编码保存
	// 将数据先保存至 buf，再持久化文件，模拟 network，也可以直接持久化至文件：NewEncoder(fi) Encode 即可；
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)

	err := enc.Encode(P{1, 2, 3, "menglima", true})
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("vt.dat", buf.Bytes(), 0644)

	// 读取解码
	fi, err := os.Open("vt.dat")
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(fi)

	q := &Q{}
	err = dec.Decode(&q)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", q)
	fmt.Println(reflect.TypeOf(q.B))
}
