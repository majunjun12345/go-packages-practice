package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	verson = "0.01"
)

/*
	go run main.go -name=masanqi
	go run main.go -name masanqi
	go run main.go -v    bool 类型不用显示赋值,调用则与默认值相反
	go run main.go -v=true   也可以
*/

var (
	ip   *int
	name string
	v    bool
	h    bool
)

func init() {
	ip = flag.Int("flagname", 12345, "help message for flagname!")
	flag.StringVar(&name, "name", "mamengli", "your full name.")

	flag.BoolVar(&v, "v", false, "show version!")
	flag.BoolVar(&h, "h", false, "show help!")
}

func main() {
	flag.Parse()

	if v {
		fmt.Printf("verson:%v\n", verson)
		os.Exit(0)
	}

	if h {
		flag.Usage()
	}

	fmt.Println(&ip, name, v, h)
	fmt.Println(flag.Args()) // 用来获取没有指定 flag 的参数
	fmt.Println(flag.Arg(0)) // 指定具体哪个

}
