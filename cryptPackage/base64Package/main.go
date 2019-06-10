package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

/*
	buf := new(bytes.Buffer) // 实现了 writer 和 reader 接口 buf.String()
*/

func main() {

	s := "www.baidu.com"

	// 自定义的 base64 编码，一般用于URL和文件名
	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))
	fmt.Println(s64)

	// 标准的
	enc_std := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(enc_std)
	dec_std, _ := base64.StdEncoding.DecodeString(enc_std)
	fmt.Println(string(dec_std))

	// 流式编码
	src := []byte("this is a test string.")
	// buf := bytes.Buffer{}  这种方式不行
	buf := new(bytes.Buffer) // 实现了 writer 和 reader 接口
	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	defer encoder.Close()
	encoder.Write(src)
	fmt.Println("stream enc:", buf.String())

	// 流式解码
	decoder := base64.NewDecoder(base64.StdEncoding, buf)
	buf_ := make([]byte, 2)
	dest := ""
	for {
		n, err := decoder.Read(buf_)
		if n == 0 || err != nil {
			break
		}
		dest += string(buf_[:n])
	}
	fmt.Println("stream de:", dest)

	// url 的，也可以用于文件名
	url := "http://www.baidu.com"
	encUrl := base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println("enc url:", encUrl)
	decUrl, _ := base64.URLEncoding.DecodeString(encUrl)
	fmt.Println("dec url:", string(decUrl))
}
