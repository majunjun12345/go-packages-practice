package main

import (
	"fmt"
	"strconv"
)

/*

	就是字符串和其他类型之间的相互转换
*/

func main() {

	// 将 所有类型 转换为字节串添加在末尾
	fmt.Println(string(strconv.AppendBool([]byte("majun:"), true))) // majun:true
	fmt.Println(string(strconv.AppendInt([]byte("majun:"), 9, 10))) // majun:9

	// 将 所有类型 转换为字符串
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatFloat(1.58765, 'f', 3, 32)) // 'e':十进制指数, 'E':二进制指数, 'f':没有指数
	fmt.Println(strconv.FormatInt(9, 10))                 // 10 表示十进制
	fmt.Println(strconv.Itoa(123))                        // 上面的简写

	/*
		将字符串转换为布尔值
		真值: 1, T, t, TRUE, true, True
		假值: 0, F, f, FALSE, false, False
	*/
	fmt.Println(strconv.ParseBool("t")) // true
	fmt.Println(strconv.ParseBool("f")) // false

	// 将字符串转换为数值
	fmt.Println(strconv.ParseInt("123", 10, 0)) // 123, 十进制, 0 for int8
	fmt.Println(strconv.Atoi("123"))            // 上面的简写

	// 使带上双引号
	fmt.Println(strconv.Quote("ma")) // "ma" 使之带上双引号

	// 判断是否为可打印字符, 空格可打印, \n 不可打印
	c := strconv.IsPrint('\u263a')
	fmt.Println(c)

	bel := strconv.IsPrint('\n')
	fmt.Println(bel)

}
