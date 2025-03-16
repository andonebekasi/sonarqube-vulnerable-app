package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("postgres", "user=admin password=admin dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	// 🚨 SQL Injection Vulnerability 🚨
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, password)
	row := db.QueryRow(query) // ⛔ Tidak menggunakan parameterized query

	var userID int
	err := row.Scan(&userID)
	if err != nil {
		fmt.Fprintf(w, "Login failed!")
	} else {
		fmt.Fprintf(w, "Welcome, %s!", username)
	}
}

func main() {
	initDB()

	http.HandleFunc("/login", handleLogin)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
