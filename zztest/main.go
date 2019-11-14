package main

import (
	"archive/tar"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

func main() {
	// test1()
	// test2()
	// test3()
	// DefaultValueOfStruct()
	// t()
	// t1()
	// ok := strings.HasSuffix("/Users/majun/sa6/test.tar", ".tar")
	// fmt.Println(ok)
	// testTar("test.tar")

	p := getCurrentDirectory()
	fmt.Println("path:", p)
}

func testTar(fpath string) {
	f, err := os.Open(fpath)
	if err != nil {
		fmt.Println("11111err:", err)
	}
	defer f.Close()
	tarRead := tar.NewReader(f)
	for {
		header, err := tarRead.Next()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("ERROR: cannot read tar file, error=[%v]\n", err)
			return
		}
		if header.FileInfo().IsDir() {
			continue
		}
		if strings.HasPrefix(filepath.Base(header.Name), ".") {
			continue
		}
	}
}

func t1() {
	fmt.Println("a" < "b")
	fmt.Println("2019-08-31" > "2019-08-31")
}

func test1() {
	state := "aHR0cDovL2xvY2FsaG9zdDo4MDgwL2YvIy90YXNr"
	url, er := base64.RawURLEncoding.DecodeString(state)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(string(url))

	t := time.Now().AddDate(0, -3, 0)
	verifyTime1 := t.Format("2006.01.02")
	fmt.Println(verifyTime1)

	var t1 int64 = 1562337405
	fmt.Println(time.Unix(t1, 0).Format("2006.01.02"))
}

func test2() {
	data, err := json.Marshal([]interface{}{"majun", "mamengli"})
	fmt.Println(data, err, len(data))

	params := []interface{}{}
	dec := json.NewDecoder(bytes.NewReader(data))
	err2 := dec.Decode(&params)
	fmt.Println(err2, params)
	fmt.Println(len(params))
}

type user struct {
	Name string
	Age  int
}

// 测试返回值为地址  error
func test3() (u *user) {
	u.Age = 19
	u.Name = "mamengli"
	return
}

// right
func test4() (u user) {
	u.Age = 19
	u.Name = "mamengli"
	return
}

type Person struct {
	Name string `defaultValue:"mengliam"`
	Age  int    `defaultValue:"21"`
}

func DefaultValueOfStruct() {
	p := Person{
		Name: "mamengli",
	}
	fmt.Printf("info:%v", p)
}

// time.Time 的零值
func t() {
	t := time.Time{}
	fmt.Println(t.IsZero())
}

func RootPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Panicln("发生错误", err.Error())
	}
	i := strings.LastIndex(s, "\\")
	path := s[0 : i+1]
	return path
}

func getCurrentPath() string {
	_, filename, _, ok := runtime.Caller(1)
	var cwdPath string
	if ok {
		cwdPath = path.Join(path.Dir(filename), "") // the the main function file directory
	} else {
		cwdPath = "./"
	}
	return cwdPath
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Debug(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
