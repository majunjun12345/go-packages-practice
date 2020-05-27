package main

import (
	"encoding/json"
	"io/ioutil"
)

func main() {
	data, err := json.MarshalIndent(Errors, "", "	")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("errors.json", data, 0644)
}
