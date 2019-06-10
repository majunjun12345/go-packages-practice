package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

/*
	sha256.New() 返回值实现了 writer 接口,向里面写入待加密的 byte 即可
	hmac.New 也一样

	io.Copy
	io.WriteString

	一般输出 16 进制字符串,通过 %x 或 hex 实现
*/

func main() {
	// code := getSha256Code("mamengli")
	// fmt.Println(code)

	hashFileCode := hashFile("Dcokerfile_bar")
	fmt.Println(hashFileCode)

	// hmac 秘钥加密
	hamcCode := getHmacCode("mamengli")
	fmt.Println(hamcCode)
}

// hash string
func getSha256Code(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	md := h.Sum(nil)                // sum 是添加额外的 byte 至 hash 头,一般没必要使用
	mdStr := hex.EncodeToString(md) // 加密之后的要以 16 进制输出,通过格式化 %x 也可以输出为 十六进制
	// return fmt.Sprintf("%x", h.Sum([]byte("ma")))
	return mdStr
}

// hash file
func hashFile(filePath string) string {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	h := sha256.New()
	io.Copy(h, f)
	return hex.EncodeToString(h.Sum(nil))
}

func getHmacCode(s string) string {
	h := hmac.New(sha256.New, []byte("myKey"))
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}
