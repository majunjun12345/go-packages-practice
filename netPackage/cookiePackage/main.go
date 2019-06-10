package main

import (
	"fmt"
	"net/http"
)

/*
	通过浏览器访问
	curl 需要带上及保存 cookie
*/

func main() {
	http.HandleFunc("/set", Set)
	http.HandleFunc("/get", Get)
	http.ListenAndServe(":8080", nil)
}

func Set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
}

func Get(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	fmt.Println("c:", c)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, c.String())
}
