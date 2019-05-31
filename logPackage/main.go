package main

import (
	"io"
	"log"
	"os"
)

/*
	log package:
	三个设置：prefix flag output，output 可以设置多个
	三种输出：
	print：就是普通 print
	panic：print + panic
	fatal：print + os.Exit(1)
*/

var logger log.Logger

func init() {
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logger.SetPrefix("【info】")
	fi, err := os.OpenFile("logtest.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	out := io.MultiWriter(fi, os.Stdout)
	logger.SetOutput(out)
}

func main() {
	logger.Println("menglima   --=")
}
