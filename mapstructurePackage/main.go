package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

/*
	map 与 struct 之间的相互转换
	struct_interface 转换为相应的 struct：struct -> interface -> struct

	encoding/json 还可以实现 json 与 map 之间的转换，不只是 struct
*/

type params struct {
	Happy  string                 `json:"happy"`
	Id     string                 `json:"id"`
	Query  []string               `json:"query"`
	Fields map[string]interface{} `json:"fields"`
	//TmpTest  string                 `json:"tmp_test"` 错误的
	//TmpTests string                 `json:"tmpTests"` 正确的
}

func main() {
	// struct2Interface2Struct()
	map2Struct()
}

//  -----------------struct -> interface -> struct
func struct2Interface2Struct() {
	p := params{
		Happy: "haha",
		Id:    "12345",
		Query: []string{"me", "ma"},
		Fields: map[string]interface{}{
			"name": "menglima",
		},
	}
	interface2Struct(p)
}

func interface2Struct(data interface{}) {
	var p params
	err := mapstructure.Decode(data, &p)
	fmt.Println(p, err)
}

// ---------------------- map to struct

func map2Struct() {
	tmp := make(map[string]interface{})
	//tmp["id"] = "123"
	//tmp["query"] = []string{"qwe", "wer", "ert"}

	param := `
			{
				"happy": "sds",
				"query":[
					"qwde",
					"wer"
				],
   				"fields":{
       			"name":"Doria",
       			"love":"HelloKitty"
   			},
				"id":"1234"
			}
`

	var confg params
	json.Unmarshal([]byte(param), &tmp)
	err := mapstructure.Decode(tmp, &confg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("get the map : ", tmp)
	fmt.Println("get the struct : ", confg)
}
