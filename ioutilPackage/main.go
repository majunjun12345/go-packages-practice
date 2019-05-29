package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// rwx 421 user group other
	err := ioutil.WriteFile("current.bak", []byte("testwrite"), 0644)
	CheckErr(err)

	/*
		读取文件
	*/
	file, err := os.Open("current.bak")
	CheckErr(err)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	/*
		读取网络内容
		网路数据记得要随时 close
	*/
	resp, err := http.Get("http://www.baidu.com")
	CheckErr(err)
	content, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(string(content))

	/*
		读取目录，不能递归
	*/
	fileinfos, err := ioutil.ReadDir("../0fragment")
	CheckErr(err)
	for _, fi := range fileinfos {
		fmt.Println(fi.Name())
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
