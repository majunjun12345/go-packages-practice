package main

import (
	"bytes"
	"log"
	"net/http"

	"github.com/gorilla/rpc/json"
)

type Args struct {
	A, B int
}

type Result int

func main() {
	url := "http://localhost:1234/rpc"

	args := &Args{
		A: 2,
		B: 3,
	}

	message, err := json.EncodeClientRequest("Arith.Multiply", args)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result Result
	err = json.DecodeClientResponse(resp.Body, &result)
	if err != nil {
		log.Fatalf("Couldn't decode response. %s", err)
	}
	log.Printf("%d*%d=%d\n", args.A, args.B, result)
}
