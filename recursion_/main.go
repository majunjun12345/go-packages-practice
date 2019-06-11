package main

import "fmt"

func main() {
	fmt.Println(fib(0))
	fmt.Println(fib(10))
}

// 0 1 1 2 3 5 8
func fib(n int) int {
	if n == 0 {
		return 0
	}
	a, b := 0, 1
	for i := 0; i < n-1; i++ {
		a, b = b, a+b
	}
	return b
}
