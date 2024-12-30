package no_graceful

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// RunNoGracefulServer starts a server without graceful shutdown
func RunNoGracefulServer() {
    // Set up the route handler for incoming requests at the root URL "/"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Processing request...")  // Log when a request is being processed
        time.Sleep(10 * time.Second)         // Simulate a long-running request (delay for 10 seconds)
        fmt.Fprintln(w, "Request completed!") // Send a response after the long task is done
    })

    // Log the start of the server without graceful shutdown
    log.Println("Starting server without graceful shutdown on :8080")
    
    // Start the HTTP server on port 8080 and wait for incoming requests
    // If there is an error (other than server shutdown), log the error and terminate the program
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
