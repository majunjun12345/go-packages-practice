package p

import (
	"fmt"
	"unsafe"
)

// 内存地址是 16进制 的

type V struct {
	i int32 // 4字节 32 位
	j int64 // 8字节 64 位
	k byte  // 1字节 8 位
}

func (v *V) Println() {
	fmt.Printf("i=%d, j=%d, k=%d \n", v.i, v.j, v.k)
}

// 定义一个打印变量地址的方法
func (v *V) PrintAddress() {
	fmt.Printf("i=%p, j=%p, k=%p \n", &v.i, &v.j, &v.k)
}

// MemoryAlignment 内存对齐
func MemoryAlignment() {
	v := &V{}
	fmt.Println(unsafe.Sizeof(v)) // 8  32位中为 4

	// 为什么会是24字节呢，int32大小为4，int64大小为8，byte大小为1，加起来是13，这就是发生了内存对齐的原因。
	/*
		下面简单详述一下go中的内存对齐：
		默认对齐值；
		32 位：4
		64 位：8

		1.对于具体类型来说，对齐值=min(编译器默认对齐值，类型大小Sizeof长度)。也就是在默认设置的对齐值和类型的内存占用大小之间，取最小值为该类型的对齐值。

		2.struct 在每个字段都内存对齐之后，其本身也要进行对齐，对齐值=min(默认对齐值，字段最大类型长度)。这条也很好理解，struct的所有字段中，最大的那个类型的长度以及默认对齐值之间，取最小的那个。简单说就是结构体整体的大小为它最大类型长度的倍数。

		第一个i占用4字节，它是结构体第一个，不需要对齐。
		第二个j占用8字节，它的对齐值通过 unsafe.Alignof(int64(0)) 函数得出也为8个字节，应用上面的第一条规则，前面只有四个字节的偏移量，4不是8的倍数，
		所以在i的后面需要填充4个字节：iiii----|jjjjjjjj。-为填充，一般是填充0。到这里整体就已经用了16个字节了。
		第三个k占用1个字节，它的对齐值通过 unsafe.Alignof(byte(0)) 函数得出也为1个字节。应用上面的第一条规则，16是1的倍数。不需要填充了。
		到这里v的字节占用变成了 iiii----|jjjjjjjj|k。接下里应用上面的第二条规则，17不是最大类型int64长度8的倍数，所以还需要填充。
		iiii----|jjjjjjjj|k------- ,在后面再填充7个字节，达到24字节后，就满足上面的第二条规则了。至此就是v为何占用24字节。
	*/
	fmt.Println(unsafe.Sizeof(*v)) // 24  32位中为 16
}

// MemoryHandle 内存操作
func MemoryHandle() {
	v := &V{}

	// 操作第一个私有变量 i，因为在go中结构体的指针就对应是第一个成员变量的指针，所以可以将 v 转换为第一个成员变量 i 的指针
	i := (*int32)(unsafe.Pointer(v)) // 先将 v 转换为普通指针，再将普通指针转换为 int32 指针
	*i = 100
	v.Println() // 解引用并赋值
	v.PrintAddress()

	// 操作第二私有个变量 j，j到i偏移了8个字节；   内存首地址                i 所占字节数              填充字节数
	j := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int32(0))) + 4)) // 第一个 4 是 i 的字节，第二个 4 是填充的 4 个字节
	*j = 200
	v.Println()
	v.PrintAddress()

	// 操作第三个变量 k，k 到 j 偏移了 16 个字节； 内存首地址              i 所占字节数              填充字节数         j 所占字节数
	k := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int32(0))) + 4 + uintptr(unsafe.Sizeof(int64(0))))) // 第一个 4 是 i 的字节，第二个 4 是填充的 4 个字节, 第三个 8 是 j 所占字节数
	*k = 5
	v.Println()
	v.PrintAddress()

	// unsafe.Sizeof 返回类型在内存中所占字节数
	fmt.Println("sizeof1:", unsafe.Sizeof(int(0)))     // 8
	fmt.Println("sizeof2:", unsafe.Sizeof(int32(0)))   // 4
	fmt.Println("sizeof3:", unsafe.Sizeof(int64(0)))   // 8
	fmt.Println("sizeof4:", unsafe.Sizeof(byte(0)))    // 1
	fmt.Println("sizeof5:", unsafe.Sizeof(string(""))) // 16

	// unsafe.Offsetof() 返回字段相对于首部的偏移量
	fmt.Println("Offsetof1:", unsafe.Offsetof(v.i)) // 0
	fmt.Println("Offsetof2:", unsafe.Offsetof(v.j)) // 8
	fmt.Println("Offsetof3:", unsafe.Offsetof(v.k)) // 16

	// uintptr 转换为可运算的指针，也就是在内存中的地址
	fmt.Println("uintptr i:", uintptr(unsafe.Pointer(v)))                                                                     // 824633811552
	fmt.Println("uintptr j:", uintptr(unsafe.Pointer(v))+uintptr(unsafe.Sizeof(int32(0)))+4)                                  // 824633811552 + 8
	fmt.Println("uintptr k:", uintptr(unsafe.Pointer(v))+uintptr(unsafe.Sizeof(int32(0)))+4+uintptr(unsafe.Sizeof(int64(0)))) // 824633811552 + 16
}
