package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testGoScripts/0INACTION/routerRedisRestful/db"
	"testGoScripts/0INACTION/routerRedisRestful/models"

	mux "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to blog service</h1>")
}

func Insert(w http.ResponseWriter, r *http.Request, _ mux.Params) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}
	fmt.Println("1111")
	if err := r.Body.Close(); err != nil { // 这里注意要关闭
		panic(err)
	}

	var u models.User
	err = json.Unmarshal(body, &u)
	if err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil { // 忘记了吗？
			panic(err)
		}
	}
	fmt.Println("2222")
	er := db.Insert(&u)
	if er != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil { // 忘记了吗？
			panic(err)
		}
	}
	fmt.Println("3333")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func Get(w http.ResponseWriter, r *http.Request, params mux.Params) {
	id, err := strconv.Atoi(params.ByName("uid"))
	if err != nil {
		panic(err)
	}
	u, err := db.FindOneByID(id)
	if err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}
