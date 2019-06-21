package main

import (
	"fmt"
	"reflect"
)

/*
	reflect 有两个数据类型: Type(数据类型) 和 Value(值)

	反射是一种检查存储在 接口变量 中的 <类型 值> 的机制
	通过 TypeOf 可以访问到 Type, 通过 ValueOf 可以访问到 Value, 从而可以进一步得到这个接口的结构类型和对其值进行操作;
*/

type User struct {
	Name    string `gogo:"name"`
	Age     int    `gogo:"int"`
	Married bool   `married:"married"`
}

func (u User) hello() {
	fmt.Printf("hello,my name is:%s", u.Name)
}

func main() {
	// var s string = "mamengli"
	// TString(s)

	u := User{"mengliam", 21, false}
	TStruct(u)
}

// 获取简单对象的类型和值
func TString(s interface{}) {
	t := reflect.TypeOf(s)  // string
	v := reflect.ValueOf(s) // mamenlgi
	fmt.Println(t, v)
}

func TStruct(s interface{}) {
	t := reflect.TypeOf(s)  // main.User
	v := reflect.ValueOf(s) // &{mengliam 21 false}

	// 获取结构体字段
	for i := 0; i < t.NumField(); i++ { // s 不能是指针
		field := t.Field(i)
		value := v.Field(i).Interface()
		/*
			Name string mengliam
			Age int 21
			Married bool false
		*/
		fmt.Println(field.Name, field.Type, value)
	}

	// 获取方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}

}
