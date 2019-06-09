package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/alecthomas/kingpin"
)

var (
	gzipServer       = kingpin.Command("gzipServer", "compress file to gzip ")
	compressToGzip   = gzipServer.Flag("compressToGzip", "compress the images into tar.gz or not").Short('c').Bool() // 压缩
	srcDir           = gzipServer.Flag("srcDir", "the direstory to gzip").Default("/Users/majun/Desktop/Secret_19050801005").Short('s').String()
	descFile         = gzipServer.Flag("descFile", "the gzip file to decompress").Default("./beeworker.tar.gz").Short('d').String()
	decompressServer = kingpin.Command("decompressServer", "decompress the gzip file")
	srcFile          = decompressServer.Flag("srcFile", "the source file to decompress").Default("./beeworker.tar.gz").Short('c').String()
	decompressDir    = decompressServer.Flag("decompressDir", "the directory gzip decompressed to").Default("./").Short('d').String()
)

func main() {
	switch kingpin.Parse() {
	case "gzipServer":
		fmt.Println(*compressToGzip)
		if *compressToGzip {
			Compress(*srcDir, *descFile)
		}
	case "decompressServer":
		decompress(*srcFile, *decompressDir)
	default:
	}
}

// 归档但不压缩
func Compress(src, desc string) {
	// 创建 tar writer文件对象
	tarFile, err := os.Create(desc)
	defer tarFile.Close()

	// gzip 压缩
	gWrite := gzip.NewWriter(tarFile)
	defer gWrite.Close()

	checkErr(err)
	tarWrite := tar.NewWriter(gWrite)

	// 如果关闭失败会造成tar包不完整，所以必须监测
	defer func() {
		err := tarWrite.Close()
		checkErr(err)
	}()

	// 写入 headerinfo 和 content 即可
	// 使用 walk 可以递归对嵌套目录进行打包
	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() { // 文件夹没有 head 和 content
			headerInfo, err := tar.FileInfoHeader(info, "")
			checkErr(err)
			headerInfo.Name, _ = filepath.Rel(filepath.Dir(src), path) // 相对路径，name 一定要改
			fileContent, err := ioutil.ReadFile(path)                  // 打开文件并读取内容，是对 os.open 的封装，os.open() 只是打开文件但还需要额外读取数据
			checkErr(err)
			err = tarWrite.WriteHeader(headerInfo)
			checkErr(err)
			tarWrite.Write(fileContent)
		}
		return nil
	})
	checkErr(err)
	fmt.Println("tar success")
}

// 解包
func decompress(src, dest string) {
	// 创建 tar reader 对象
	fi, err := os.Open(src)
	checkErr(err)

	gfi, err := gzip.NewReader(fi)
	checkErr(err)
	defer gfi.Close()

	tarRead := tar.NewReader(gfi)

	// 遍历获取 tar 包文件
	for headInfo, err := tarRead.Next(); err != io.EOF; headInfo, err = tarRead.Next() {
		checkErr(err)

		// 创建目录层级
		_, err = os.Stat(filepath.Join(dest, filepath.Dir(headInfo.Name))) // 判断文件夹是否存在
		if os.IsNotExist(err) {
			os.MkdirAll(filepath.Join(dest, filepath.Dir(headInfo.Name)), os.ModePerm) // os.ModeDir 权限
		}

		// 创建文件并写进去即可
		f, err := os.Create(filepath.Join(dest, headInfo.Name))
		defer f.Close()
		checkErr(err)
		_, err = io.Copy(f, tarRead)
		checkErr(err)
	}
	fmt.Println("untar success")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
