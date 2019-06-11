package main

var p *int // *int 是指向 int 类型的指针，其零值为 nil，也就是 p　是一个指针类型

func main() {
	/*
		给指针类型赋值有两种方式
	*/
	// one
	*p = 15 // **int　表示地址上的值

	// b
	b := 13
	p = &b // &　是取 址 符号
}
