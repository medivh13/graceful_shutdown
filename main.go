package main

import (
	gr "graceful_shutdown/graceful"
	noGr "graceful_shutdown/no_graceful"
	"flag"
	"log"
)

func main() {
	// Define a flag to select the server type
	serverType := flag.String("server", "graceful", "Choose server type: 'graceful' or 'no-graceful'")
	flag.Parse()

	switch *serverType {
	case "no-graceful":
		log.Println("Running server without graceful shutdown...")
		noGr.RunNoGracefulServer()
	case "graceful":
		log.Println("Running server with graceful shutdown...")
		gr.RunGracefulServer()
	default:
		log.Fatalf("Unknown server type: %s. Use 'graceful' or 'no-graceful'", *serverType)
	}
}
