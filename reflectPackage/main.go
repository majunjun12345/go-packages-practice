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

func (u *User) Hello() string {
	return fmt.Sprintf("hello,my name is:%s\n", u.Name)
}

func main() {

	u := &User{"mengliam", 21, false}
	TStruct(u)

	// main1()
}

func TStruct(s interface{}) {
	// Value、Type、kind
	// Type 表示 interface{} 的实际类型，Kind 表示特别类型
	v := reflect.ValueOf(s) // &{mengliam 21 false}
	t := reflect.TypeOf(s)  // main.User
	k := t.Kind()           // 如果传入的是指针: ptr，如果传入的是结构体: struct
	fmt.Println("typeof valueof kind", t, v, k)

	// 获取结构体字段
	switch reflect.TypeOf(s).Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			/*
				Name string mengliam
				Age int 21
				Married bool false
			*/
			fmt.Println(t.Field(i).Name, t.Field(i).Type, v.Field(i).Interface())

			switch v.Field(i).Kind() {
			case reflect.String:
				fmt.Println(v.Field(i).String())
			case reflect.Int:
				fmt.Println(v.Field(i).Int())
			case reflect.Bool:
				fmt.Println(v.Field(i).Bool())
			default:
				fmt.Println("Unsupported type", t.Field(i).Name, t.Field(i).Type, v.Field(i).Interface())
			}
		}
	case reflect.Ptr: // 如果是指针，则要剥离指针
		fmt.Println("is ptr")
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.String:
				fmt.Println(v.Field(i).String())
			case reflect.Int:
				fmt.Println(v.Field(i).Int())
			case reflect.Bool:
				fmt.Println(v.Field(i).Bool())
			default:
				fmt.Println("Unsupported type", t.Field(i).Name, t.Field(i).Type, v.Field(i).Interface())
			}
		}
	}
	// Int 和 String 可以帮助我们分别取出 reflect.Value 作为 int64 和 string
	{
		a := 56
		x := reflect.ValueOf(a).Int()
		fmt.Printf("type:%T value:%v\n", x, x) // type:int64 value:56
		b := "Naveen"
		y := reflect.ValueOf(b).String()
		fmt.Printf("type:%T value:%v\n", y, y) // type:string value:Naveen
	}

	// 获取方法， 方法必须是可导出的
	{
		fmt.Println(reflect.TypeOf(s).NumMethod(), reflect.TypeOf(s).Name())
		for i := 0; i < reflect.TypeOf(s).NumMethod(); i++ {
			m := t.Method(i)
			fmt.Println("=======", m.Name)
		}

		m := v.MethodByName("Hello")
		/*
			这里会 panic，因为上面将传入的 v 被 剥离了指针
			User类型是*User的基底类型
			在Go的指针知识中，有一条规则：一个指针类型拥有它以及它的基底类型为接收者类型的所有方法，而它的基底类型却只能拥有以它本身为接收者类型的方法。
		*/
		rs := m.Call(nil)
		fmt.Println("result:", rs[0].Interface().(string))
	}
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
