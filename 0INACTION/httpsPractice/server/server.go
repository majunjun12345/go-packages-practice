package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello https"))
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServeTLS(":443", "../keys/ca/server.crt", "../keys/ca/server.key", nil)
	if err != nil {
		fmt.Println(err)
	}
}
