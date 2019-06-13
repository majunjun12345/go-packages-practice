package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	filepathPackage()
}

// filepath
func filepathPackage() {
	absPath, _ := filepath.Abs("main.go") // 返回所给文件的绝对路径 /Users/majun/go/src/testGoScripts/filepathPackage/main.go 加上当前脚本的执行路径
	fmt.Println(absPath)

	basePath := filepath.Base(absPath + "/") // 返回路径的最后一个文件名或夹，可以兼容最后面有 /
	fmt.Println(basePath)

	cleanPath := filepath.Clean(absPath) // 不知道是干啥
	fmt.Println(cleanPath)

	dir := filepath.Dir(absPath + "/") // 返回最后一个文件所在目录，不兼容 /
	fmt.Println(dir)

	files, _ := filepath.Glob("./*.go") // 返回指定目录下所有匹配的文件，当前目录下
	fmt.Println(files)

	fmt.Println(filepath.IsAbs(absPath)) // 查看路径是否是绝对路径
	fmt.Println(filepath.IsAbs(basePath))

	p := filepath.Join("a", "b") // a/b
	fmt.Println(p)

	fmt.Println(filepath.Match("*", "haha")) // 不知道这是在干啥

}
