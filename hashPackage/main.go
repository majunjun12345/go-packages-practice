package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	// m := md5String()
	// fmt.Println(m)

	result, _ := hashFile()
	fmt.Println(result)
}

// 同 sha1
func md5String() string {
	m := md5.New() // 返回的是 hash.hash
	m.Write([]byte("haha"))
	return hex.EncodeToString(m.Sum([]byte(""))) // sum 加盐
}

// 获取文件的 sha256 值
func hashFile() (string, error) {
	fi, err := os.Open("test.txt")
	if err != nil {
		return "", nil
	}
	m := md5.New()

	_, err = io.Copy(m, fi)
	if err != nil {
		return "", nil
	}
	value := hex.EncodeToString(m.Sum([]byte("")))
	return value, nil
}
