package main

import (
	"encoding/base64"
	"fmt"
	"time"
)

func main() {
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
