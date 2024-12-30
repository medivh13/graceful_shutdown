package graceful

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// RunGracefulServer starts an HTTP server and handles graceful shutdown
func RunGracefulServer() {
	// Create a new HTTP server with a simple handler
	server := &http.Server{
		Addr: ":8080", // Set the address and port for the server
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Processing request...")  // Log when a request is being processed
			time.Sleep(10 * time.Second)          // Simulate a long-running request
			fmt.Fprintln(w, "Request completed!") // Respond once the long task is finished
		}),
	}

	// Run the server in a separate goroutine to avoid blocking the main thread
	go func() {
		log.Println("Starting server with graceful shutdown on :8080")
		// Listen and serve, handle any error except for the expected "server closed" error
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Create a channel to listen for system interrupt signals (e.g., Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt) // Notify the channel when an interrupt signal is received

	// Wait for the interrupt signal to be received (blocking)
	<-stop
	log.Println("Shutting down gracefully...") // Log when the shutdown process starts

	// Create a context with a timeout to limit how long the server can take to shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure the cancel function is called when the shutdown is complete

	// Attempt to gracefully shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		// If there is an error during shutdown, log it and terminate the program
		log.Fatalf("Graceful shutdown failed: %v", err)
	}
	log.Println("Server stopped") // Log when the server has successfully stopped
}
