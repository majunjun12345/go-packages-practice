package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Info struct {
	Metadata    map[string]interface{}
	Count       int    `json:"count"`
	Description string `json:"description"`
	*Person
	*Info
}

// 将结构体(有嵌套结构体)通过 reflect 的方式转换为 map
func main() {
	info := &Info{
		Metadata:    map[string]interface{}{"addr": "127.0.0.1"},
		Count:       2,
		Description: "information",
	}
	info.Person = &Person{
		Name: "mamengli",
		Age:  19,
	}
	fmt.Println("=======", Struct2Map(info))
}

func Struct2Map(info interface{}) (kv map[string]interface{}) {

	t := reflect.TypeOf(info)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}

	v := reflect.Indirect(reflect.ValueOf(info)) // 等同于剥离指针

	kv = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {

		if t.Field(i).Type.Kind() == reflect.Ptr {
			if v.Field(i).IsNil() {
				continue
			}
			fmt.Println(v.Field(i).Interface())
			for k, v := range Struct2Map(v.Field(i).Interface()) {
				kv[k] = v
			}
		} else {
			kv[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	return
}
