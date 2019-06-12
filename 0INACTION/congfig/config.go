package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"gopkg.in/gcfg.v1"
)

type Configuration struct {
	Enabled bool   `json:"enabled" xml:"enabled" ini:"enabled" yaml:"enabled"`
	Path    string `json:"path" xml:"path" ini:"path" yaml:"path"`
}

var conf Configuration

type IniConfiguration struct {
	Section Configuration `ini:"section"`
}

func main() {
	// JsonConf()
	// XmlConf()
	// IniConf()
	YamlConf()
}

// yaml
func YamlConf() {
	b, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", conf)
}

// ini
func IniConf() {
	c := &IniConfiguration{}
	err := gcfg.ReadFileInto(c, "conf.ini")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", c)
}

// xml
func XmlConf() {
	fi, err := os.Open("conf.xml")
	defer fi.Close()
	if err != nil {
		panic(err)
	}
	dec := xml.NewDecoder(fi)
	err = dec.Decode(&conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", conf)
}

// json
func JsonConf() {
	fi, err := os.Open("conf.json")
	defer fi.Close()
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(fi)
	err = dec.Decode(&conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", conf)
}
