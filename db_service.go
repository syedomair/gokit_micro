package main

import (
	"database/sql"
	"github.com/go-kit/kit/log"
	_ "github.com/lib/pq"
	"os"
)

func DBService(logger log.Logger) (*sql.DB, error) {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Log("err", "The data source arguments are not valid")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Log("err", "Database connection error")
		return nil, err
	}
	return db, err
}
