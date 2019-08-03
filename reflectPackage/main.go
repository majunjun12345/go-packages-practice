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

	// u := User{"mengliam", 21, false}
	// TStruct(u)

	main1()
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

	m := v.MethodByName("hello")
	params := make([]reflect.Value, 1)
	fmt.Println(m.Call(params))
}

// 通过反射调用结构体的方法
func main1() {
	//通过反射的方式调用结构体类型的方法
	var setNameStr string = "SetName"
	var addAgeStr string = "AddAge"
	user := UserInfo{
		Id:   1,
		Name: "env107",
		Age:  18,
	}
	//1.获取到结构体类型变量的反射类型
	refUser := reflect.ValueOf(&user) //需要传入指针，后面再解析
	fmt.Println(refUser)
	//2.获取确切的方法名
	//带参数调用方式
	setNameMethod := refUser.MethodByName(setNameStr)
	args := []reflect.Value{reflect.ValueOf("Mike")} //构造一个类型为reflect.Value的切片
	setNameMethod.Call(args)                         //返回Value类型
	//不带参数调用方式
	addAgeMethod := refUser.MethodByName(addAgeStr)
	addAgeMethod.Call(make([]reflect.Value, 0))

	fmt.Println("User.Name = ", user.Name)
	fmt.Println("User.Age = ", user.Age)

}

type UserInfo struct {
	Id   int
	Name string
	Age  int
}

//ToString方法
func (u UserInfo) String() string {
	return "User[ Id " + string(u.Id) + "]"
}

//设置Name方法
func (u *UserInfo) SetName(name string) string {
	oldName := u.Name
	u.Name = name
	return oldName
}

//年龄数+1
func (u *UserInfo) AddAge() bool {
	u.Age++
	return true
}

//测试方法
func (u UserInfo) TestUser() {
	fmt.Println("我只是输出某些内容而已....")
}
