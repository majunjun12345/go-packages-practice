package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/get/1", g1)
	http.HandleFunc("/post/1", p1)
	http.HandleFunc("/post/2", p3)

	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	http.ListenAndServe(":1234", nil)
}

/*
	r.ParseForm()             需要解析后才能获取参数
	r.Form.Get("age")  21
	r.Form["age"]      [21]

	r.URL.Query().Get("age")  可以直接获取参数

	r.FormValue("age")        可以直接获取参数 get/post
	r.PostFormValue("name")   post
*/
func g1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("age"))      // 21
	fmt.Println(r.PostFormValue("name")) // post 请求的参数

	r.ParseForm()
	fmt.Println(r.Form.Get("age")) // 21
	fmt.Println(r.Form["age"])     // [21]

	fmt.Println(r.URL.Query().Get("name"))
	fmt.Println(r.URL.Query().Get("age"))

	fmt.Println(r.URL.RequestURI())
	fmt.Println(r.Header.Get("name"))
	w.Write([]byte("success"))
	return
}

func p1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("name"))
	fmt.Println(r.PostFormValue("name"))

	w.Write([]byte("success post"))
	return
}

func p3(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form.Get("city"))

	buf := make([]byte, 1024)
	n, err := r.Body.Read(buf)
	fmt.Println(n, err, buf)
	w.Write([]byte("success post2"))
	return
}
