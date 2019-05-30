package main

import (
	"fmt"
	"os"
)

/*
	os.File 的一系列操作
	os.create 打开的文件可读可写，没有就创建
	os.open 打开的文件只读，如果是目录，获取的是 os.FileInfo，不会递归操作

	seek 用来设置文件的偏移量，read 操作后偏移量会改变，ReadAt 不会改变偏移量
*/

func main() {

	fileObj, err := os.Open("test.txt")
	CheckErr(err)

	content := make([]byte, 10)
	n, err := fileObj.Read(content)
	CheckErr(err)
	fmt.Println(n, string(content))

	n, err = fileObj.ReadAt(content, 10) // read 偏移量以后的数据
	fmt.Println(string(content))

	fileObj.Read(content)
	fmt.Println(string(content))

	curOffset, err := fileObj.Seek(0, os.SEEK_CUR)
	fmt.Println("curOffset", curOffset)

	/*
		只有 open 是一个目录时,才会有 Readdir 方法
	*/
	fileObj2, err := os.Open("../../ioPackage")
	fileinfos, err := fileObj2.Readdir(2)
	CheckErr(err)
	for _, findo := range fileinfos {
		fmt.Println(findo.Name())
	}

	names, err := fileObj2.Readdirnames(0) // 有问题，没读全
	CheckErr(err)
	fmt.Println("names:", names)

	/*
		0 表示从文件头开始偏移,1 表示从当前位置开始偏移, 2 表示从文件末尾开始偏移
		读取的是偏移量 15 后的数据

		read 之后，偏移量向后移动 len(content)
	*/
	n1, err := fileObj.Seek(15, 0)
	CheckErr(err)
	fmt.Println(n1)
	fileObj.Read(content)
	fmt.Println(string(content))

	curOffset, err4 := fileObj.Seek(0, os.SEEK_CUR) // 获取当前的偏移量
	CheckErr(err4)
	fmt.Println(curOffset)

	fileInfo, _ := fileObj.Stat()
	fmt.Printf("%#v", fileInfo)

	// n3, err := fileObj.Write([]byte("mamengli"))   // 报错
	// CheckErr(err)
	// fmt.Println("n3", n3)

	fileObj.Close() // 记得随时关闭文件
	fileObj2.Close()

	fileObjC, err := os.Create("create.txt")
	CheckErr(err)
	nc, err := fileObjC.WriteString("mamengli")
	CheckErr(err)
	fmt.Println(nc)
	CheckErr(fileObjC.Sync()) // sync 就是 bufio 包的 flush，确保数据写入磁盘
	fileObjC.Close()

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
