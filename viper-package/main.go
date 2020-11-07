package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

/*
[Go Viper – 初探](https://www.dazhuanlan.com/2019/12/12/5df18474d9d5d/)

Viper是配置文件的解决方案。支持：
	设置默认值
	可以读取如下格式的配置文件：JSON、TOML、YAML、HCL
	监控配置文件改动，并热加载配置文件
	从环境变量读取配置
	从远程配置中心读取配置（etcd/consul），并监控变动
	从命令行 flag 读取配置
	从缓存中读取配置
	支持直接设置配置项的值

Viper读取配置顺序
	viper.Set()所设置的值
	命令行 flag
	环境变量
	配置文件
	配置中心：etcd/consul
	默认值
*/

func main() {
	// ReadFromEnv()
	ReadConfigOnUpdata()
}

// 环境变量中读取配置文件：
func ReadFromEnv() {
	viper.AutomaticEnv()
	fmt.Println(viper.Get("PATH"))
}

func ReadConfigOnUpdata() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	// 全局的
	viper.SetConfigName("config")
	// viper.AddConfigPath(".")
	viper.SetConfigFile("./conf/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s n", err))
	}

	// 实时重新读取配置文件
	viper.WatchConfig()
	// 如果配置文件发生改变就会触发
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println(viper.Get("denyu"))
	})

	fmt.Println(viper.Get("denyu"))

	waitGroup.Wait()
}

// ----------------------------------------------------------------

var (
	cfg = pflag.StringP("config", "c", "conf/config.yaml", "config file")
)

func main1() {
	pflag.Parse()

	// init flag
	if err := Init(*cfg); err != nil {
		panic(err)
	}

	for {
		fmt.Println("runmode:", viper.Get("runmode"))
		fmt.Println("PATH:", viper.Get("PATH"))
		time.Sleep(4 * time.Second)
	}
}

type Config struct {
	Name string
}

func Init(cfgName string) error {
	c := &Config{
		Name: cfgName,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}
	// 监控配置文件变化并热加载程序
	c.watchConfig()
	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 指定配置文件，
	} else {
		viper.AddConfigPath("conf")   // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config") // 不包含扩展名
		// 这里也可以从配置中心中加载配置
	}
	viper.SetConfigType("yml") // 设置配置文件格式为 YAML
	viper.AutomaticEnv()       // 读取环境变量
	// viper.SetEnvPrefix("APISERVER")           // 读取环境变量的前缀为 APISERVER
	// replacer := strings.NewReplacer(".", "_") // APISERVER_LOG_PATH => log.path
	// viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil { // viper 解析配置文件
		return err
	}
	return nil
}

// 监控配置文件变化并加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 这里要进行程序的重启
		log.Printf("Config file changed: %s", e.Name)
	})
}
