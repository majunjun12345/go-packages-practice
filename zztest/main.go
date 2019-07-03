package main

import (
	"fmt"
	"time"
)

func main() {
	// state := "aHR0cDovL2xvY2FsaG9zdDo4MDgwL2YvIy90YXNr"
	// url, er := base64.RawURLEncoding.DecodeString(state)
	// if er != nil {
	// 	fmt.Println(er)
	// }
	// fmt.Println(string(url))
	t := time.Now()
	verifyTime1 := t.Format("2006.01.02")
	fmt.Println(verifyTime1)
}
