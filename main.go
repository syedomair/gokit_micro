package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
)

type Env struct {
	db *sql.DB
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	db, _ := DBService()
	env := &Env{db: db}
	http.Handle("/public/books", booksIndex(env))
	port := os.Getenv("PORT")
	logger.Log("msg", "HTTP", "addr", ":"+port)
	logger.Log("err", http.ListenAndServe(":"+port, nil))
}

func booksIndex(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		bks, _ := PublicBooks(env.db)
		js, err := json.Marshal(bks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

func PublicBooks(db *sql.DB) ([]BookResponse, error) {

	rows, err := db.Query("SELECT book.id, " +
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
		fmt.Println("No Results Found")
	}
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	books := make([]BookResponse, 0)
	for rows.Next() {
		book := BookResponse{}
		err := rows.Scan(&book.Id, &book.UserId, &book.FirstName, &book.LastName, &book.Name, &book.Description, &book.Publish)
		if err != nil {
			fmt.Println(err)
		}
		books = append(books, book)
	}
	return books, nil
}
