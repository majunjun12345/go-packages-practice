package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	// Get1()
	// Post1()
	// Post2()
	// Post3()
	p4()
}

// get
/*
	add 增加参数值
	set 覆盖已有参数值
	req.URL.RawQuery = q.Encode() encode 后才能正常传值
*/
func Get1() {
	client := http.Client{}

	req, err := http.NewRequest("GET", "http://127.0.0.1:1234/get/1", nil)
	CheckErr(err)

	q := req.URL.Query()
	q.Add("name", "majun")
	q.Add("name", "menglimamama")
	q.Set("age", "21")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("name", "menglima")
	resp, err := client.Do(req)
	CheckErr(err)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	fmt.Println(string(content))
}

// post
/*
	post 请求需增加请求头 Content-Type
	二进制流
*/
func Post1() {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	// 方式 一
	// req, err := http.NewRequest("POST", "http://127.0.0.1:1234/post/1", strings.NewReader("name=mamengli"))
	// CheckErr(err)
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// resp, err := client.Do(req)

	// 方式 二
	resp, err := client.Post("http://127.0.0.1:1234/post/1", "application/x-www-form-urlencoded;charset=UTF-8", strings.NewReader("name=mamengli"))

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	fmt.Println(string(content))
}

// PostForm 这种方式不需要加上述请求头
func Post2() {
	resp, err := http.PostForm("http://127.0.0.1:1234/post/1", url.Values{"name": {"masanqi"}})
	CheckErr(err)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	fmt.Println(string(content))
}

func Post3() {
	resp, err := http.Post("http://127.0.0.1:1234/post/2?city=wuhan", "multipart/form-data", bytes.NewReader([]byte("hahahhahaaaahaaaaa")))
	CheckErr(err)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	fmt.Println(string(content))
}

func p4() {
	u := fmt.Sprintf("http://localhost:8081/file/mpupload/uppart?username=%s&uploadid=%s&index=%d", "userName", "uploadId", "i")
	resp, err := http.Post(u, "multipart/form-data", bytes.NewReader([]byte("hahahhahaaaahaaaaa"))) // 6. post 的使用
	CheckErr(err)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	fmt.Println(string(content))
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
