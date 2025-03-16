package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    server := &http.Server{
        Addr:           ":8080",
        Handler:        http.HandlerFunc(handler),
        ReadTimeout:    10 * time.Second, // Timeout baca
        WriteTimeout:   10 * time.Second, // Timeout tulis
        IdleTimeout:    120 * time.Second, // Timeout idle
    }

    // Menangani error jika server gagal dimulai
    err := server.ListenAndServe()
    if err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}
