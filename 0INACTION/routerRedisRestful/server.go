package main

import (
	"log"
	"net/http"
	"testGoScripts/0INACTION/routerRedisRestful/routes"
)

func main() {
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8002", r))
}
