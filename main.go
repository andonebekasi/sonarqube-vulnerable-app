package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/vuln", func(w http.ResponseWriter, r *http.Request) {
		userInput := r.URL.Query().Get("id")
		db, err := sql.Open("postgres", "postgres://user:pass@localhost/dbname?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		// SQL Injection vulnerability
		query := "SELECT * FROM users WHERE id = " + userInput
		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, "Error executing query", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		fmt.Fprintf(w, "Query executed: %s", query)
	})

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
