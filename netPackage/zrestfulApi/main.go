package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	json NewEncoder 可以向 writer 中写入 经过 encode 转换为 json 的 struct
	json NewDecoder 向 reader 中读取数据，并经过 decode 后，json 可以转换为 struct
	流式传输，适合网络中数据的传输

	NewEncoder 只是向 w 中写入，还未返回，后面还可以加个 w.write(例如：delete 方法)

	omitempty：为空则不输出
*/

type Person struct {
	ID        string  `json:"id, omitempty"`
	FirstName string  `json:"firstname, omitempty"`
	LastName  string  `json:"lastname, omitempty"`
	Address   Address `json:"address, omitempty"`
}

type Address struct {
	City     string `json:"city, omitempty"`
	Province string `json:"province"`
}

var people []Person

func main() {
	people = append(people, Person{ID: "1", FirstName: "ma", LastName: "mengli", Address: Address{City: "wuhan", Province: "hubei"}})
	people = append(people, Person{ID: "2", FirstName: "ma", LastName: "sanqi", Address: Address{City: "xiantao", Province: "hubei"}})
	router := mux.NewRouter()
	router.HandleFunc("/get/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/gets", GetPeople).Methods("GET")
	router.HandleFunc("/add", AddPerson).Methods("POST")
	router.HandleFunc("/delete/{id}", DeletePerson).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	if id, ok := mux.Vars(r)["id"]; ok {
		for i, p := range people {
			if p.ID == id {
				people = append(people[:i], people[i+1:]...)
				json.NewEncoder(w).Encode(people)
			} else {
				w.Write([]byte("no such person"))
			}
		}
	}
}

func AddPerson(w http.ResponseWriter, r *http.Request) {
	var p Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		panic(nil)
	}
	people = append(people, p)
	json.NewEncoder(w).Encode(people)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	people, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	w.Write(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if ok {
		for _, p := range people {
			if p.ID == id {
				json.NewEncoder(w).Encode(p)
				// person, err := json.Marshal(p)
				// if err != nil {
				// 	panic(err)
				// }
				// w.Write(person)
			}
		}
	}
}
