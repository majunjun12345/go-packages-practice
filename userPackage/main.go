package main

import (
	"fmt"
	"os/user"
)

func main() {
	u, _ := user.Current()

	fmt.Println(u.HomeDir)
}
