package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"testGoScripts/cmdArgsPackage/cobraPackage/cmd/server1"
	"testGoScripts/cmdArgsPackage/cobraPackage/cmd/server2"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	author  string
)

var rootCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Run the gRPC hello-world server",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	// 以时间戳为全局随机种子
	rand.Seed(time.Now().UnixNano())
	cobra.OnInitialize(initConfig)

	// 所有服务通用的变量
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yml", "config file (default is config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&author, "author", "a", "menglima", "Author name for copyright attribution")

	// 自动绑定所有命令行参数，如果只需要其中某个，可以用viper.BingPflag()选择性绑定
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlags(rootCmd.PersistentFlags())

	// 添加两个子命令，分别对应独立的服务
	rootCmd.AddCommand(server1.Server1Cmd)
	rootCmd.AddCommand(server2.Server2Cmd)
}

// 这里可以初始化全局的配置
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		fmt.Println("initConfigFrom cfgFile: ", cfgFile)
	} else {
		// 也可以从配置中心读取配置
	}

	// 从环境变量中读取，k8s 普及后用的越来越多
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 热加载程序
	})
}
