package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetItems",
		"GET",
		"/items",
		GetItems,
	},
	Route{
		"GetItem",
		"GET",
		"/items/{id}",
		GetItem,
	},
	Route{
		"CreateItem",
		"POST",
		"/items/{todoId}",
		CreateItem,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
