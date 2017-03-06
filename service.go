package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    //"os"
    "fmt"
)

type BookResponse struct {
    Id              int64  `json:"id"`
    UserId          int64  `json:"user_id"`
    FirstName       string `json:"first_name"`
    LastName        string `json:"last_name"`
    Name            string  `json:"book_name" `
    Description     string  `json:"description" `
    Publish         bool  `json:"publish" `
} 


type BookService interface {
	PublicBooks() []BookResponse
}

type bookService struct{}

func (bookService) PublicBooks() []BookResponse {

        //db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
        dbUrl := "postgres://lrfegrijyhscgm:460c27d5d24257af04c92fcb2062a17affacc538875471c499cdf274b596d739@ec2-54-225-99-171.compute-1.amazonaws.com:5432/d3fptb61ed7ps2"
        db, err := sql.Open("postgres", dbUrl)
        if err != nil {
            fmt.Println("Error: The data source arguments are not valid")
        }
        err = db.Ping()
        if err != nil {
            fmt.Println("Error: Could not establish a connection with the database")
        }
        rows, err := db.Query("SELECT book.id, book.user_id, u.first_name, u.last_name, book.name, book.description, book.publish FROM book join public.user as u on book.user_id = u.id  WHERE book.publish = true  ")
        if err == sql.ErrNoRows {
            fmt.Println("No Results Found")
        }
        if err != nil {
            fmt.Println(err)
        }
        defer rows.Close()
        bks := make([]BookResponse, 0)
        for rows.Next() {
            bk := BookResponse{}
            err := rows.Scan(&bk.Id, &bk.UserId, &bk.FirstName, &bk.LastName, &bk.Name, &bk.Description, &bk.Publish)
            if err != nil {
                fmt.Println(err)
            }
            bks = append(bks, bk)
        }
        fmt.Println(bks)
       for _, bk := range bks {
           fmt.Printf("%s, %s, %s, %s, %s, %s, %s \n", bk.Id, bk.UserId, bk.FirstName, bk.LastName, bk.Name, bk.Description, bk.Publish)
       }
	return bks
}
