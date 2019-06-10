package main

import "fmt"

type Mutatable struct {
	a int
	b int
}

func (m *Mutatable) pTest() {
	m.a = 5
	m.b = 7
}

func (m Mutatable) Test() {
	m.a = 3
	m.b = 4
}

// 值方法和指针方法
func pointer_value_method() {
	m := Mutatable{}
	fmt.Printf("origin:%+v\n", m)

	m.Test()
	fmt.Printf("origin:%+v\n", m)

	m.pTest()
	fmt.Printf("origin:%+v\n", m)

	m1 := &Mutatable{1, 2}
	fmt.Printf("origin:%+v\n", m1)
	m1.Test()
	fmt.Printf("origin:%+v\n", m1)
	m1.pTest()
	fmt.Printf("origin:%+v\n", m1)
}

func main() {
	// pointer_value_method()

	// 继承 依旧可以调用嵌套字段的方法
	s := Student{Human{"mamengli", 19, "155"}, "free"}
	s.SayHi()

	// 重写, 调用的是重写的方法
	h := Employee{Human{"masanqi", 21, "180"}, "ks"}
	h.SayHi()
}

// ------------------------
// method 继承
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名字段
	school string
}

// 在 Human 上定义一个方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, my name is %s, %d years old,my phonenumber is %s\n", h.name, h.age, h.phone)
}

// ----------------------重写
type Employee struct {
	Human
	company string
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}
