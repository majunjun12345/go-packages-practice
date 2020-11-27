package main

import (
	"fmt"
	"reflect"

	"github.com/unknwon/com"
)

/*
	类型转换
*/

func main() {
	state := "1"
	result := com.StrTo(state).MustInt()
	fmt.Println(reflect.TypeOf(result))

	i := 5
	s := com.ToStr(i)
	fmt.Println(reflect.TypeOf(s))
}
