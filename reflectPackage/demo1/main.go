package main

import (
	"fmt"
	"reflect"
)

/*
	[Go 系列教程 —— 34. 反射](https://studygolang.com/articles/13178)

	任何一变量的底层是由 类型 + 值 组成, reflect 包能够识别 interface{} 底层的 类型 和 值;
*/

type Person struct {
	Name string
	Age  int
	Sex  string
}

type tool struct {
	cap string
	key string
}

func (t *tool) print() {
	fmt.Println(t.cap, t.key)
}

func (p Person) Say(msg string) {
	fmt.Println("hello,", msg)
}

func (p Person) PrintInfo(t *tool) {
	t.cap = "green"
	t.key = "long"
	fmt.Printf("姓名:%s, 年龄:%d, 性别:%s, 参数tool内容:%s %s\n", p.Name, p.Age, p.Sex, t.key, t.cap)
}

type service struct {
	servers map[string]reflect.Method
	rcvr    reflect.Value
	typ     reflect.Type
}

func MakeService(rep interface{}) *service {
	ser := service{}
	ser.typ = reflect.TypeOf(rep)
	ser.rcvr = reflect.ValueOf(rep)
	// name返回其包中的类型名称，举个例子，这里会返回Person，tool
	name := reflect.Indirect(ser.rcvr).Type().Name()
	fmt.Println(name)
	ser.servers = map[string]reflect.Method{}
	fmt.Println(ser.typ.NumMethod(), ser.typ.Name())
	for i := 0; i < reflect.TypeOf(rep).NumMethod(); i++ {
		method := ser.typ.Method(i)
		// mtype := method.Type
		//mtype := method.Type	// reflect.method
		mname := method.Name // string
		fmt.Println("mname : ", mname)
		ser.servers[mname] = method
	}
	return &ser
}

func main() {
	p1 := Person{"Rbuy", 20, "男"}
	// 得到这个对象的全部方法，string对应reflect.method
	methods := MakeService(p1)
	// 利用得到的methods来调用其值
	methname := "PrintInfo"
	if method, ok := methods.servers[methname]; ok {
		// 得到第一个此method第1参数的Type，第零个当然就是结构体本身了
		replyType := method.Type.In(1)
		replyType = replyType.Elem() // Elem会返回对
		// New returns a Value representing a pointer to a new zero value for the specified type.
		replyv := reflect.New(replyType)
		function := method.Func
		function.Call([]reflect.Value{methods.rcvr, replyv})
		// 此时我们已经拿到了返回值
	}
}
