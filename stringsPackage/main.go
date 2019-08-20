package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	BasicUse()
}

func BasicUse() {

	// join
	new_s := strings.Join([]string{"zhangsan", "lisi"}, ".")
	fmt.Println("=====:", new_s)

	// to
	s := "MengLiMa"
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToTitle(s))

	// compare
	var s1 string = "Welcome to The WORld of go!"
	var s2 string = "Welcome go The WORld of go!"
	fmt.Println(strings.Compare(s1, s2))   // 0 -1 1
	fmt.Println(strings.EqualFold(s1, s2)) // false

	// trim
	var s3 = "Goodbye, world!"
	fmt.Println(strings.Trim(s3, "!"))          // 去除首尾指定字符 还有 right left
	fmt.Println(strings.TrimSpace(s3))          // 去除首尾空格
	fmt.Println(strings.TrimPrefix(s3, "Good")) // 去除头不指定字符串，还有 sufix

	// split
	s4 := "a:b:c:d:e-f-g-h:b:c:d:"
	fmt.Println(strings.Split(s4, "-"))      // [a:b:c:d:e f g h:b:c:d:]
	fmt.Println(strings.Split(s4, "b"))      // [a: :c:d:e-f-g-h: :c:d:]
	fmt.Println(strings.SplitN(s4, "-", 2))  // [a:b:c:d:e f-g-h:b:c:d:] 2 是指两个部分，切割一次
	fmt.Println(strings.SplitAfter(s4, "-")) // [a:b:c:d:e- f- g- h:b:c:d:]  保留了 sep

	// method
	s5 := "hjhbknmk,ddhbfe,f"
	fmt.Println(strings.HasPrefix(s5, "h")) // 首部是否包含 还有 sufix
	fmt.Println(strings.Contains(s5, "dd"))
	fmt.Println(strings.Index(s5, "k"))     // 返回字符索引
	fmt.Println(strings.LastIndex(s5, "f")) // 最后一个 字符 的索引

	// replace
	s6 := "arfhesmksv mde"
	fmt.Println(strings.Replace(s6, "s", "t", 1)) // 最后 n 表示替换次数，0 表示不替换
	fmt.Println(strings.Replace(s6, "s", "t", 0))
	fmt.Println(strings.ReplaceAll(s6, "s", "t")) // 全部替换

	// io
	s7 := "ajerbfneinowernwigni"
	r := strings.NewReader(s7)
	// s := "MengLiMa"
	r.Reset(s) // 清空所有缓存，并且将缓存读切换到 s

	for {
		b, err := r.ReadByte()
		fmt.Println(r.Size()) // size 一直不变
		fmt.Println(r.Len())  // len 随着 read 逐渐变少
		// r.Seek(10, 0)         // 将 read 的指针 挪到 指定处
		fmt.Println(string(b))
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			panic(err)
		}
	}
}
