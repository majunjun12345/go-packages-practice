package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	// test1()
	// test2()
	test3()
}

func test1() {
	state := "aHR0cDovL2xvY2FsaG9zdDo4MDgwL2YvIy90YXNr"
	url, er := base64.RawURLEncoding.DecodeString(state)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(string(url))

	t := time.Now().AddDate(0, -3, 0)
	verifyTime1 := t.Format("2006.01.02")
	fmt.Println(verifyTime1)

	var t1 int64 = 1562337405
	fmt.Println(time.Unix(t1, 0).Format("2006.01.02"))
}

func test2() {
	data, err := json.Marshal([]interface{}{"majun", "mamengli"})
	fmt.Println(data, err, len(data))

	params := []interface{}{}
	dec := json.NewDecoder(bytes.NewReader(data))
	err2 := dec.Decode(&params)
	fmt.Println(err2, params)
	fmt.Println(len(params))
}

type user struct {
	Name string
	Age  int
}

// 测试返回值为地址  error
func test3() (u *user) {
	u.Age = 19
	u.Name = "mamengli"
	return
}

// right
func test4() (u user) {
	u.Age = 19
	u.Name = "mamengli"
	return
}
