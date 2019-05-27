package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
)

/*
0. mux 与 DefaultServeMux
	http.NewServeMux()：
	使用该方法后，后续可以通过 mux.Handle 或 mux.HandlerFunc 实现注册 api 路由
	作为 ListenAndServe 的第二个参数
	DefaultServeMux：
	使用 http.Handle 或 http.HandlerFunc 注册路由
	不必传递 ListenAndServe 的第二个参数

1. HandlerFunc、Handle、Handler
	HandlerFunc:
	http.HandlerFunc 会将自定义的 函数，转换为实现了 ServerHTTP 方法的 handler；
	ServerHTTP 内部将会调用 函数方法；  ServerHTTP 内部是 api 的真正执行内容，中间件同理；
	同 mux.HandlerFunc;
	Handle:
	Handle 接收的第二个参数则必须实现 ServerHTTP 方法，必须是 http.handler
	handler:
	接口，实现了 ServerHTTP 的所有类型都是 handler

2. 写中间件的两种方式
	- 结构体
		必须有个字段是 http.handler，表示 api handler
		该结构体实现 ServerHTTP 方法，方法内逻辑是中间件的实现目的
		必须显示调用 http.handler 对应的 ServerHTTP 方法；
			理由：
			在没有中间件的情况下，net/http 框架会找到 api handler 并调用其 ServerHTTP 方法 ...
			加上中间件后，net/http 框架会首先找到中间件的 ServerHTTP 方法，如果不显示调用 api handler，那么其将不会执行
		该结构体实例化后，只能作为 mux 传参 ListenAndServe，那么对所有的 api都将会起作用；

	- 将 api handler 包裹的函数

3. 中间件与 api handler 的执行
	顺序：完全取决于 handler.ServeHTTP(w, r) 的位置; 如果多个 mid 嵌套，最先执行最外层 mid；
	完成：handler 执行完成之后就返回给客户端了
	注意：如果 mid 需要改变 response(header status body)，无论 handler.ServeHTTP(w, r) 放在哪都不保险
		可以参考：https://www.cnblogs.com/hitandrew/p/5820677.html
*/

type SingleHost struct {
	next      http.Handler
	allowHost string
}

func (sh *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		sh.next.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

// 还可以传递参数
func logMidware(next http.Handler, s string) http.Handler {
	outFunc := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("梦里马"))
		log.Println("log middleware test")
		next.ServeHTTP(w, r)
		log.Println(s)
		w.Write([]byte("mamenlgi"))
	}
	return http.HandlerFunc(outFunc)
}

func edit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("edit"))
}

// 通过 httptest 方法对响应内容进行修改
func editMid(next http.Handler, s string) http.Handler {
	myFunc := func(w http.ResponseWriter, r *http.Request) {
		rec := httptest.NewRecorder()
		next.ServeHTTP(rec, r)

		for k, v := range rec.Header() {
			w.Header()[k] = v
		}
		w.Header().Set("X-We-Modified-This", "Ma")
		w.WriteHeader(418)
		data := []byte(s)

		clen, _ := strconv.Atoi(r.Header.Get("Content-Length"))
		clen += len(data)
		r.Header.Set("Content-Length", strconv.Itoa(clen))

		w.Write(data)
		w.Write(rec.Body.Bytes())
	}
	return http.HandlerFunc(myFunc)
}

func main() {

	// sh := &SingleHost{
	// 	handler:   http.HandlerFunc(hello),
	// 	allowHost: "example.con",
	// }
	http.HandleFunc("/hello", hello)

	http.Handle("/test", logMidware(http.HandlerFunc(test), "mamengli"))
	http.Handle("/edit", editMid(http.HandlerFunc(edit), "sanqi"))

	// log.Fatal(http.ListenAndServe(":3031", sh))
	log.Fatal(http.ListenAndServe(":3031", nil))

}
