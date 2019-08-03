package main

import (
	"fmt"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

/*
	最快的 json 解释器
*/

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Birth   string `json:"births"`
	Married bool   `json:"married"`
}

func main() {
	// testMarshal()

	// testUnmarshal()

	// testGet()

	testDecode()
}

// marshal
func testMarshal() {
	b := time.Now()
	t := b.Format("2006-09-01 15:31:21")

	p := &Person{
		Name:    "mamengli",
		Age:     21,
		Birth:   t,
		Married: false,
	}

	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	result, err := jsonIterator.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}

// unmarshal
func testUnmarshal() {
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	jsonContent := []byte(`[{"name":"mamengli","age":21,"births":"2019-09-07 20:87:77","married":false},{"name":"masanqi","age":22,"births":"2019-09-07 20:87:77","married":false}]`)
	p := &[]Person{}
	jsonIterator.Unmarshal(jsonContent, p)
	fmt.Printf("%+v", p)
}

// Get
func testGet() {
	jsonContent := []byte(`{"name":"mamengli","age":21,"births":"2019-09-07 20:87:77","married":false}`)
	str := jsoniter.Get(jsonContent, "name").ToString()
	fmt.Println(str)
}

// decode
func testDecode() {
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	reader := strings.NewReader(`{"name":"mamengli","age":21,"births":"2019-09-07 20:87:77","married":false}`)
	decoder := jsonIterator.NewDecoder(reader)

	p := &Person{}
	err := decoder.Decode(p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", p)
}
