package main

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HandlerServer
}

type Routes []Route

var routes = Routes{
	Route{
		"PublicBooks",
		"GET",
		"/public/books",
		PublicBooks,
	},
}
