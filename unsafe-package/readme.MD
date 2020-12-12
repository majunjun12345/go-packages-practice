### unsafe
c 语言中提供了直接通过指针运算操作内存的方式，其他语言基本屏蔽了这点；
go 中可以使用 unsafe 包绕过 go 的内存安全机制，直接对内存进行读写；

当处理系统调用时，Go 的结构体必须和 C 的结构体拥有相同的内存结构，这时你可能除了使用 unsafe 以外，别无选择。

#### 一、内存对齐
[【转】关于C语言中数据结构的内存对齐问题](http://blog.chinaunix.net/uid-21457204-id-3061765.html)
默认对齐值；
	32 位：4
	64 位：8

为什么要内存对齐？
提高访问速度

#### 二、unsafe.Pointer
- type Pointer 
普通指针
unsafe包提供的一个类型，任何指针类型都可以转换为Pointer类型,
Pointer类型也可以转换为任何指针类型。go中的指针不能进行运算(比如:*int),只能对数据进行读取和赋值。

- uintptr 
go的一个数据类型，它不是unsafe包提供的，它是一个足够大的整数类型，可以容乃任何的指针。
上面的 Pointer类型可以转换为 uintptr类型，反过来uintptr类型也可以转换为Pointer类型。
uintptr类型才能进行指针运算。

unsafe.Pointer 可以让你无视 Go 的类型系统，完成任何类型与内建的 uintptr 类型之间的转化。
根据文档，unsafe.Pointer 可以实现四种其他类型不能的操作：
  任何类型的指针都可以被转化为Pointer
  Pointer可以被转化为任何类型的指针
  uintptr可以被转化为Pointer
  Pointer可以被转化为uintptr