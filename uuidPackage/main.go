package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	// 创建
	u1, err := uuid.NewV4() // 24a7ccb4-c5f8-4e60-a540-382daa3c4f3e
	if err != nil {
		panic(err)
	}
	fmt.Println(u1)

	// 解析
	u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}
