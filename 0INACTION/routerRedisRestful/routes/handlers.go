package routes

import (
	"fmt"
	"net/http"

	mux "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to blog service</h1>")
}
