package common

import "reflect"

func Struct2Map(obj interface{}) map[string]interface{} { // obj 不能是地址
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
