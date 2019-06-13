package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

/*
	basic authentication
	假设客户端与服务端的连接是安全的

	相当于明文传输(base64 编码), 后端校验
*/

func main() {
	http.HandleFunc("/auth", Auth)
	http.ListenAndServe(":8000", nil)
}

func Auth(w http.ResponseWriter, r *http.Request) {

	if check := CheckAuth(w, r); check {
		w.Write([]byte("hello"))
		return
	}
	w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
	w.WriteHeader(401)
	w.Write([]byte("401 Unauthorized!\n"))
}

func CheckAuth(w http.ResponseWriter, r *http.Request) bool {
	fmt.Println(r.Header.Get("Authorization"))                 // Basic dXNlcjpwYXNz
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2) // [Basic dXNlcjpwYXNz]

	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		panic(err)
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	fmt.Println("pair:", pair) // [user pass]
	if len(pair) != 2 {
		return false
	}
	return pair[0] == "user" && pair[1] == "pass"
}
