package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

func main() {
	// filepath
	// filepathPackage()

	// path
	// pathPackage()

	//url
	urlPath()
}

func urlPath() {
	example := "https://root:123456@www.baidu.com:0000/login?name=xiaoming&name=xiaoqing&age=24&age1=23#fffffff"
	urlPase, _ := url.Parse(example)
	//                        /login www.baidu.com:0000 www.baidu.com      true           /login            0000     map[name:[xiaoming xiaoqing] age:[24] age1:[23]]  name=xiaoming&name=xiaoqing&age=24&age1=23
	fmt.Println(urlPase.EscapedPath(), urlPase.Host, urlPase.Hostname(), urlPase.IsAbs(), urlPase.Path, urlPase.Port(), urlPase.Query(), urlPase.RawPath, urlPase.RawQuery)
	// /login?name=xiaoming&name=xiaoqing&age=24&age1=23   https
	fmt.Println(urlPase.RequestURI(), urlPase.Scheme)
}

// path
func pathPackage() {
	basePath := path.Base("a/b/c/") // 返回最后一个元素, 兼容最后的 /   c
	fmt.Println(basePath)

	dirPath := path.Dir("a/b/c/") // 返回最后一个元素的目录, 不兼容 /
	fmt.Println(dirPath)

	extention := path.Ext("/a/b/c/d.txt") // 返回文件扩展名 .txt
	fmt.Println(extention)

	// path.IsAbs()   判断是否是绝对路径
	// path.Join      连接路径
	// path.Split     切分目录与文件
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

	relPath, _ := filepath.Rel("/a/b/c", "/a/b/c/d/e") // d/e,后面对于前面的相对路径,同相对同绝对
	fmt.Println(relPath)

	dirname, filename := filepath.Split("/a/b/c/d.txt") // /a/b/c  d.txt  切割目录和文件名
	fmt.Println(dirname, filename)

	paths := filepath.SplitList("/a/b/c/d.txt") // [/a/b/c/d.txt] 不知道干啥
	fmt.Println(paths)

	fmt.Println(filepath.VolumeName("[/a/b/c/d.txt]")) // 返回分区名, 作用于 windows

	err := filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path) // path 是相对于当前执行脚本的路径
		if path == "../.git/logs/refs/heads/master" {
			return errors.New("out")
		}
		return nil
	})
	fmt.Println(err)
}
