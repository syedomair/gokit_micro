package main

import (
	"context"
	"database/sql"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
)

type Env struct {
	db     *sql.DB
	logger log.Logger
	ctx    context.Context
}

func main() {

	logger := log.NewLogfmtLogger(log.StdlibWriter{})
	db, err := DBService(logger)
	if err != nil {
		return
	}

	env := &Env{db: db,
		logger: logger,
		ctx:    context.Background()}

	router := NewRouter(env)

	port := os.Getenv("PORT")
	env.logger.Log("msg", "Listening at HTTP", "PORT", port)
	env.logger.Log("err", http.ListenAndServe(":"+port, router))

}
