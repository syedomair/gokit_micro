package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
)

var PublicBooks = func(env *Env) *httptransport.Server {
	var service BookService
	service = bookService{}

	publicBookHandler := httptransport.NewServer(
		env.ctx,
		makePublicBooksEndpoint(service, env),
		decodeRequest,
		encodeResponse,
	)
	return publicBookHandler
}
