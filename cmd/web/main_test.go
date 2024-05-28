package main

import (
	"net/http"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	go main() // Start the server in a goroutine

	// Give the server a second to start
	time.Sleep(1 * time.Second)

	// Make a request to our server with the / route
	resp, err := http.Get("http://localhost:4000/")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// We should get a 200 status
	if resp.StatusCode != 202 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}
