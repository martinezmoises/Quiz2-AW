package main

import (
	"fmt"
	"log"
	"net/http"
)

//HTTP server that contains multiple routes. When a request is made to a route, print a log line

func main() {
	// HandleFunc registers a route and its corresponding handler function.
	// This function handles the "/hello" route.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Log a message when the /hello route is accessed
		log.Println("Handling /hello")

		// Write "Hello, World!" to the HTTP response
		fmt.Fprintln(w, "Hello, World!")
	})

	// This function handles the "/goodbye" route.
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		// Log a message when the /goodbye route is accessed
		log.Println("Handling /goodbye")

		// Write "Goodbye, World!" to the HTTP response
		fmt.Fprintln(w, "Goodbye, World!")
	})

	// Log a message indicating that the server is starting
	log.Println("Server starting on port 3000...")

	// ListenAndServe starts an HTTP server on port 8080.
	// The second argument (nil) means we are using the default HTTP multiplexer
	// (http.DefaultServeMux), which automatically handles the registered routes.
	http.ListenAndServe(":3000", nil)
}
