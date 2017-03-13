package main

import (
	"database/sql"
	_ "github.com/lib/pq"
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
	PublicBooks(*Env) []BookResponse
}

type bookService struct{}

func (bookService) PublicBooks(env *Env) []BookResponse {

	rows, err := env.db.Query("SELECT book.id, " +
		" book.user_id, " +
		" u.first_name, " +
		" u.last_name, " +
		" book.name, " +
		" book.description, " +
		" book.publish " +
		" FROM book join public.user as u " +
		" on book.user_id = u.id  " +
		" WHERE book.publish = true  ")
	if err == sql.ErrNoRows {
		env.logger.Log("msg", "No Results Found")
	}
	if err != nil {
		env.logger.Log("err", err)
	}
	defer rows.Close()
	books := make([]BookResponse, 0)
	for rows.Next() {
		book := BookResponse{}
		err := rows.Scan(&book.Id, &book.UserId, &book.FirstName, &book.LastName, &book.Name, &book.Description, &book.Publish)
		if err != nil {
			env.logger.Log("err", err)
		}
		books = append(books, book)
	}
	return books
}
