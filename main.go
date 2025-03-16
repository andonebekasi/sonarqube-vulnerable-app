package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Potensi SQL Injection karena input langsung digunakan dalam query
	query := r.URL.Query().Get("user")
	fmt.Fprintf(w, "Hello, %s", query) // Tidak ada validasi input
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
