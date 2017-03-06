package main

import (
    "context"
    "encoding/json"
    "net/http"
    "github.com/go-kit/kit/endpoint"
)

func makePublicBooksEndpoint(bookService BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        return bookService.PublicBooks(), nil
    }
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
    return "", nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    return json.NewEncoder(w).Encode(response)
}
