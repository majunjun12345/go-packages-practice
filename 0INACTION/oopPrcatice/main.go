package main

import "fmt"

type Animal struct { // 基础结构
	Name string
	mean bool // 私有属性
}

func (a *Animal) PerformNosize(strength int, sound string) {
	if a.mean {
		strength = strength * 3
	}

	for voice := 0; voice < strength; voice++ {
		fmt.Println(sound)
	}
	fmt.Println("end")
}

type Cat struct {
	Basics       Animal // go 中使用组合的方式替代继承
	MeowStrength int
}
type Dog struct {
	Animal       // 匿名结构
	BarkStrength int
}

func (d *Dog) MakeNosize() {
	// barkStrength := d.BarkStrength
	// if d.mean {
	// 	barkStrength = barkStrength * 3
	// }

	// for bark := 0; bark < barkStrength; bark++ {
	// 	fmt.Println("bark")
	// }
	// fmt.Println("end")
	d.PerformNosize(d.BarkStrength, "bark")
}

func (c *Cat) MakeNosize() {
	// // meowStrength := c.MeowStrength
	// // if c.Basics.mean {
	// // 	meowStrength = meowStrength * 3
	// // }
	// // for meow := 0; meow < meowStrength; meow++ {
	// // 	fmt.Println("MEOW")
	// // }
	// fmt.Println("end")
	c.Basics.PerformNosize(c.MeowStrength, "MEOW")
}

// 接口声明一系列的行为,而类型则实现这种行为
type AnimalSounder interface { // 当接口中只包含一种方法时,接口使用 er 后缀命名
	MakeNosize()
}

/*
	利用接口创建多态
	只要类型实现了接口定义的行为,那么他就可以表示这个接口类型
*/
func MakeSomeNosize(animalSounder AnimalSounder) {
	animalSounder.MakeNosize()
}

func main() {
	c := &Cat{
		Basics: Animal{
			Name: "haha",
			mean: true,
		},
		MeowStrength: 3,
	}

	d := &Dog{
		Animal{
			Name: "rfsd",
			mean: true,
		},
		3, // 匿名结构实例化时,不能写字段名
	}

	MakeSomeNosize(c)
	MakeSomeNosize(d)
}
