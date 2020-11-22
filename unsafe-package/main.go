package main

import (
	"fmt"
	"unsafe"
)

/*
	不同类型的指针变换
*/

func main() {
	u := uint32(32)
	i := int32(1)
	fmt.Println(&u, &i)
	p := &i
	// (*int32)(&u): cannot convert &u (type *uint32) to type *int32
	p = (*int32)(unsafe.Pointer(&u))
	fmt.Println(p, *p)
}
