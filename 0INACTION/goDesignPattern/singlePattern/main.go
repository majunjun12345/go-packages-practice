package main

import "sync"

func main() {

}

type Config struct {
	addr string
}

var ConfigInfo *Config = nil

// 方式一, 非线程安全
func SinglePattern1() *Config {
	if ConfigInfo == nil {
		ConfigInfo = &Config{
			addr: "127.0.0.1:80",
		}
	}
	return ConfigInfo
}

// 方式二, 线程安全
func SinglePattern2() *Config { // 一般函数名是 New
	var once sync.Once
	once.Do(func() {
		ConfigInfo = &Config{
			addr: "127.0.0.1:80",
		}
	})
	return ConfigInfo
}
