package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Create a handler for the /hello route
	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Write "Hello, World!" as the response
		fmt.Fprintln(w, "Hello, World!")
	})

	// Create a handler for the /goodbye route
	goodbyeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Write "Goodbye, World!" as the response
		fmt.Fprintln(w, "Goodbye, World!")
	})

	// Create a new ServeMux (HTTP request multiplexer) to register the routes
	mux := http.NewServeMux()

	// Register the /hello route and associate it with the helloHandler
	mux.Handle("/hello", helloHandler)

	// Register the /goodbye route and associate it with the goodbyeHandler
	mux.Handle("/goodbye", goodbyeHandler)

	// Wrap the mux (which has all the registered routes) with the logging middleware.
	// This ensures every request goes through the loggingMiddleware first.
	loggedMux := loggingMiddleware(mux)

	// Start the HTTP server on port 8080, using the loggedMux (which applies middleware).
	log.Println("Server starting on port 3000...")
	http.ListenAndServe(":3000", loggedMux)
}
