package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

type HandlerServer func(*Env) *httptransport.Server

//No need to define func (f HandlerServer) ServeHTTP because it is already there in httptransport.Server

func NewRouter(env *Env) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc(env)
		handler = Logger(handler, route.Name, env)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
