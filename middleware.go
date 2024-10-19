package main

import (
	"log"
	"net/http"
	"time"
)

// loggingMiddleware is a middleware function that takes an http.Handler,
// wraps it, and logs information about each request.

// http.HandlerFunc which is being used in the place of func(ResponseWriter, *Request)
func loggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // returns a new HTTP handler.

		//middleware logic starts:

		// Start timer to measure how long the request takes
		start := time.Now()

		// Log the HTTP method and path of the incoming request
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		// Call the next handler in the chain (in this case, the actual route handler)
		next.ServeHTTP(w, r)

		// After the handler finishes, log how long the request took
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
