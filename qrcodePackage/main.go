package main

import (
	"fmt"
	"os"

	"github.com/skip2/go-qrcode"
	r "github.com/tuotoo/qrcode"
)

func main() {

	// CreateQRcode()

	Rqrcode()
}

// 生成二维码图片
func CreateQRcode() {
	err := qrcode.WriteFile("https://www.jianshu.com/writer#/notebooks/36302655/notes/45716733", qrcode.Medium, 256, "./jianshu_qrcode.png")
	if err != nil {
		panic(err)
	}
}

// 识别二维码图片
func Rqrcode() {
	fi, err := os.Open("./jianshu_qrcode.png")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	qrmatrix, err := r.Decode(fi)
	if err != nil {
		panic(err)
	}
	fmt.Println(qrmatrix.Content)
}
