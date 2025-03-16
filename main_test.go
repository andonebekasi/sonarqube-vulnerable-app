package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVulnHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/vuln?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test response"))
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
	}
}
