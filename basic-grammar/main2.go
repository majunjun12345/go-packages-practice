package main

import (
	"fmt"
)

type Action struct {
	Name string
}

var actionFactors map[string]actionBuilder

type actionBuilder func(entityType, entityID string) *Action

func main() {
	Registered(
		"menglima",
		func(entityType, entityID string) *Action {
			return &Action{
				Name: entityID,
			}
		},
	)

	fmt.Println(actionFactors["menglima"]("ma", "qwe"))
}

func Registered(name string, ac actionBuilder) {
	actionFactors[name] = ac
}

func init() {
	actionFactors = make(map[string]actionBuilder)
}
