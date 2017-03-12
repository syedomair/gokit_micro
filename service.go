package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type BookResponse struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Name        string `json:"book_name" `
	Description string `json:"description" `
	Publish     bool   `json:"publish" `
}

type BookService interface {
	PublicBooks() []BookResponse
}

type bookService struct{}

func DBService() (*sql.DB, error) {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Error: The data source arguments are not valid")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error: Database connection error")
	}
	return db, nil
}
