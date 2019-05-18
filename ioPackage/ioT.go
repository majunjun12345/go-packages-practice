package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	// srcF, err := os.Open("Dockerfile")
	// check(err)
	// descF, err := os.Create("Dcokerfile_bar")
	// check(err)

	// io.Copy(descF, srcF)

	resp, err := http.Get("http://www.baidu.com")
	check(err)
	// buf := bufio.NewReaderSize(resp, 1024)

	descF, err := os.Create("baidu")
	check(err)

	io.Copy(descF, resp.Body)

}

func check(err error) {
	if err != nil {
		panic(nil)
	}
}
