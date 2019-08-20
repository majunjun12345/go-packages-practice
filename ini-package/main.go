package main

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	appMode        string
	data           string
	protocol       string
	http_port      int
	enforce_domain bool
)

func main() {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		log.Fatal("load 'app.ini' failed:%v\n", err)
	}

	// 没有 section 就用空字符串代替
	appMode = cfg.Section("").Key("app_mode").String()

	// section 指定分区
	data = cfg.Section("paths").Key("data").String()

	// 选择限制
	protocol = cfg.Section("server").Key("protocol").In("http", []string{"http", "https"})

	// 类型转换
	http_port = cfg.Section("server").Key("http_port").MustInt(9999)
	enforce_domain = cfg.Section("server").Key("enforce_domain").MustBool(false)
}
