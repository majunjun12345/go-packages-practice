package server

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	ServerPort  string
	CertName    string
	CertPemPath string
	CertKeyPath string
)

func Serve1() (err error) {
	log.Println(ServerPort)

	// 貌似大小写不敏感
	fmt.Println("path:", viper.Get("path"))
	fmt.Println("NAME:", viper.Get("NAME"))
	fmt.Println("LISTEN:", viper.Get("listen"))
	// 嵌套
	fmt.Println("db name:", viper.Get("db.name"))

	return nil
}

func Serve2() {
	fmt.Println("this is server2")
}
