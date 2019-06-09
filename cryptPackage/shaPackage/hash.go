package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	code := getSha256Code("mamengli")
	fmt.Println(code)
}

func getSha256Code(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum([]byte("ma")))
}
