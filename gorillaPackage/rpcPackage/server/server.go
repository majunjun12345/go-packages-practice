package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Args struct {
	A, B int
}

type Arith int

type Result int

func (a *Arith) Multiply(r *http.Request, args *Args, result *Result) error {
	log.Printf("Multiplying %d with %d\n", args.A, args.B)
	*result = Result(args.A * args.B)
	return nil
}

func main() {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	arith := new(Arith)
	server.RegisterService(arith, "")

	r := mux.NewRouter()
	r.Handle("/rpc", server)
	http.ListenAndServe(":1234", r)
}
