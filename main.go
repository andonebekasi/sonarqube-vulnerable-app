package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

func vulnerableHandler(w http.ResponseWriter, r *http.Request) {
	// SQL Injection Vulnerability
	userInput := r.URL.Query().Get("user")
	query := "SELECT * FROM users WHERE name = '" + userInput + "'" // Rentan SQLi

	fmt.Fprintf(w, "Query executed: %s", query)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", vulnerableHandler)

	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
