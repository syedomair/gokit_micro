package main

import (
    "context"
    "net/http"
    "os"
    "github.com/go-kit/kit/log"
    httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
    ctx := context.Background()
    logger := log.NewLogfmtLogger(os.Stderr)

    var service BookService
    service = bookService{}

    publicBookHandler := httptransport.NewServer(
        ctx,
        makePublicBooksEndpoint(service),
        decodeRequest,
        encodeResponse,
    )

    http.Handle("/public/books", publicBookHandler)
    port := os.Getenv("PORT")
    logger.Log("msg", "HTTP", "addr", ":"+port)
    logger.Log("err", http.ListenAndServe(":"+port, nil))
}
