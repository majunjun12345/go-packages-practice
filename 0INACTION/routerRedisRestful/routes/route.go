package routes

import (
	mux "github.com/julienschmidt/httprouter"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  mux.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		Index,
	},
	Route{
		"index",
		"POST",
		"/user/insert",
		Insert,
	},
	Route{
		"index",
		"GET",
		"/user/:uid",
		Get,
	},
}

func NewRouter() *mux.Router {

	router := mux.New()

	for _, route := range routes {

		router.Handle(route.Method, route.Pattern, route.Handle)
	}

	return router
}
