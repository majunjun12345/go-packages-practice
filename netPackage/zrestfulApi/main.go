package main

import (
	"github.com/gorilla/mux"
)

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
people = append(people, Person{ID:1, FirstName:"ma", LastName:"mengli", Address:Address{City:"wuhan", Province:"hubei"}})
people = append(people, Person{ID:2, FirstName:"ma", LastName:"sanqi", Address:Address{City:"xiantao", Province:"hubei"}})


func main() {
	router := mux.NewRouter()
}
