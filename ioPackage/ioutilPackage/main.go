package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	// rwx 421 user group other
	err := ioutil.WriteFile("current.bak", []byte("testwrite"), 0644)
	CheckErr(err)

	/*
		直接读取文件内容,参数是文件名
	*/
	// ioutil.ReadAll(reader)
	content2, err := ioutil.ReadFile("current.bak")
	CheckErr(err)
	fmt.Println(string(content2))

	/*
		读取目录,但是不能递归
	*/
	fileinfos, err := ioutil.ReadDir("../")
	CheckErr(err)
	for _, fi := range fileinfos {
		fmt.Println(fi.Name())
	}

	/*
		创建临时目录, "" 表示 /tmp
	*/
	dirname, err := ioutil.TempDir("", "example")
	CheckErr(err)
	fmt.Println(dirname) // /tmp/example331421827
	err = ioutil.WriteFile(filepath.Join(dirname, "test.txt"), []byte("memgnlima\n"), 0644)
	CheckErr(err)

	/*
		创建临时文件
	*/
	tmpfile, err := ioutil.TempFile("", "file.txt")
	fmt.Println(tmpfile.Name())
	_, err = tmpfile.Write([]byte("menglima"))
	CheckErr(err)

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
