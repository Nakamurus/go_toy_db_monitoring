package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		connStr := "postgres://dbuser:secret@db:5432/test?sslmode=disable"

		db, err := sql.Open("postgres", connStr)

		if err != nil {
			http.Error(w, "Cannot connect to database", http.StatusInternalServerError)
			return
		}

		defer db.Close()

		err = db.Ping()
		if err != nil {
			http.Error(w, "Cannot ping database", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("DB is healthy!"))
	})

	http.ListenAndServe(":8080", nil)
}
