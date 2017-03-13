package main

import (
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string, env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		env.logger.Log(
			"Method", r.Method,
			"URI", r.RequestURI,
			"MethodName", name,
			"Time", time.Since(start))
	})
}
