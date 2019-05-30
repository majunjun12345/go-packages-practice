package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	state := "aHR0cDovL2xvY2FsaG9zdDo4MDgwL2YvIy90YXNr"
	url, er := base64.RawURLEncoding.DecodeString(state)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(string(url))
}
