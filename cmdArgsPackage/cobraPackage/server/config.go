package server

import (
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Name   string `mapstructure:"NAME"`
	Listen int    `mapstructure:"LISTEN"`
	GoPath string `mapstructure:"GOPATH"`
}

func GetConf() (*ServerConfig, error) {
	var conf ServerConfig
	err := viper.Unmarshal(&conf)
	if err != nil {
		return nil, multierror.Prefix(err, "invalid values of viper")
	}
	return &conf, nil
}
