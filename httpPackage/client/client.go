package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"
)

/*
	[[译] 初学者需要了解的Go语言中的HTTP timeout](https://studygolang.com/articles/26359?fr=sidebar)
*/

func main() {
	// Get1()
	// Post1()
	// Post2()
	// Post3()
	p4()
}

func newHttpClient() *http.Client {

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100 // 连接池大小，为 0 表示没有限制，断开超过部分
	/*
		默认 MaxConnsPerHost=2
		当一个 host 请求多次时，服务端同时只能处理两次，其他请求将处于 TIME_WAIT 状态，将会消耗服务器资源，直至崩溃
		所以这里最好和 MaxIdleConns 保持一致，为 100
	*/
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	client := &http.Client{
		/*
			超时
			完整的请求周期：dialer(三次握手)、TLS 握手、请求头和请求体的生成和发送、响应头及响应体的接收;
			client 处的 timeout 定义完整的请求周期超时时间，这里最好不超过 10s;
		*/
		Timeout: time.Second * 10,
		/*
			transport的主要功能其实就是缓存了长连接，用于大量http请求场景下的连接复用，减少发送请求时TCP(TLS)连接建立的时间损耗，
			同时transport还能对连接做一些限制，如连接超时时间，每个host的最大连接数等；

			DialContext：用于控制 dial 即三次握手的超时时间；
			KeepAlive：心跳间隔；
		*/
		Transport: t,
	}
	return client
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

func p5() {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Hour)
	}))
	defer svr.Close()

	client := &http.Client{
		// Timeout: time.Second * 5, // 如果这里没有 timeout，将会一直等待一小时
	}
	fmt.Println("making request")
	_, err := client.Get(svr.URL)
	fmt.Println("finished request", err)
}
