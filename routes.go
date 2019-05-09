package main

import (
	"net/http"
)

// Route represents route entity.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of `Route`.
type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"TodoIndex", "GET", "/todos", TodoIndex},
	Route{"TodoShow", "GET", "/todos/{id}", TodoShow},
	Route{"TodoCreate", "POST", "/todos", TodoCreate},
}
