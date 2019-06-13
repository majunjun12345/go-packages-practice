package main

import (
	"bytes"
	"io/ioutil"

	"github.com/disintegration/imaging"
)

func main() {
	ImageResize("menglima.JPG", "2.jpg")
}

func ImageResize(filename, targetFilename string) {
	imgData, err := ioutil.ReadFile(filename)
	CheckErr(err)
	buf := bytes.NewBuffer(imgData)
	img, err := imaging.Decode(buf)
	CheckErr(err)
	// newImg := imaging.Resize(img, 300, 0, imaging.Lanczos)  // 为 0 则表示等比例
	newImg := imaging.Fill(img, 100, 100, imaging.Center, imaging.Lanczos) // 将会按比例裁剪，以中心为基准点
	imaging.Save(newImg, targetFilename)
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
