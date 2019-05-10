package main

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	dir := "/Users/majun/go/src/tmp/"
	dest := "/Users/majun/go/src/test_script"
	Compress(dir, dest)
}

func Compress(dir, dest string) {

	zipPath := filepath.Join(filepath.Dir(dest), "zipTest.zip")
	zFile, err := os.Create(zipPath)
	if err != nil {
		panic(err)
	}
	wZip := zip.NewWriter(zFile)
	defer wZip.Close()

	f, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range f {
		if file.IsDir() {
			continue
		}
		fileName := file.Name()
		if fileName == ".DS_Store" {
			continue
		}
		fWZip, err := wZip.Create(fileName)
		fileContent, err := ioutil.ReadFile(filepath.Join(dir, fileName))

		if err != nil {
			panic(err)
		}
		_, err = fWZip.Write(fileContent)
		if err != nil {
			panic(err)
		}
	}
}
