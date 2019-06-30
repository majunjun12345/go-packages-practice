package main

import "fmt"

func init() {
	fmt.Println("init in z.go")
}

func z() int64 {
	fmt.Println("calling z() in z.go")
	return 3
}

var _ int64 = z()
