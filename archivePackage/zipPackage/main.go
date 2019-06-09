package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// err := UnZip("manageTool.zip", "./")
	// ChectErr(err)
	ToZip()
}

/*
	以 fi.Mode() 模式新建文件，不然显示没权限
*/
func UnZip(archieveFilePath, target string) error {
	err := os.MkdirAll(target, 0755)
	if err != nil {
		return err
	}
	archievrFiles, err := zip.OpenReader(archieveFilePath)
	defer archievrFiles.Close()

	for _, fi := range archievrFiles.File {
		if fi.FileInfo().IsDir() {
			os.MkdirAll(filepath.Join(target, fi.Name), fi.Mode())
			continue
		}

		fileContent, err := fi.Open()
		defer fileContent.Close()
		if err != nil {
			return err
		}

		targetFilePath := filepath.Join(target, fi.Name)
		targetFile, err := os.OpenFile(targetFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, fi.Mode())
		defer targetFile.Close()
		if err != nil {
			return err
		}

		_, err = io.Copy(targetFile, fileContent)
		if err != nil {
			return err
		}
	}
	return nil
}

func ToZip() {
	// buf := new(bytes.Buffer)
	// buf.
	// w := zip.NewWriter(buf)
	fi, err := os.OpenFile("test.tar", os.O_CREATE|os.O_WRONLY, 0666)
	defer fi.Close()
	ChectErr(err)
	w := zip.NewWriter(fi)
	defer w.Close()

	err1 := filepath.Walk("./manageTool", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			panic(err)
			return err
		}
		f, err := w.Create(path)

		if err != nil {
			panic(err)
		}
		content, err := ioutil.ReadFile(path)
		_, err = f.Write(content)
		if err != nil {
			panic(err)
		}
		return err
	})
	ChectErr(err1)
}

func ChectErr(err error) {
	if err != nil {
		panic(err)
	}
}
