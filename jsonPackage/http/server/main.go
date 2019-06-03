package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// curl -i -H "Content-Type: application/json" -X POST  --data '{"Id":"1", "Balance":50}' http://localhost:8080/

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// 解析请求
		var u User
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Println(u.Id)

		// 构造响应
		u1 := User{Id: "US123", Balance: 8}
		json.NewEncoder(w).Encode(u1) // 直接在 w 里面写，参数是个 writer 对象就行
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	Id      string
	Balance uint64
}
