// HTTP Server
// - Using `net/http` package for HTTP server tasks, requests/responses, etc.

// Incoming requests take a path and then a handler function

package main

import (
	"fmt"
	"net/http"
)

func sendHelloResponse(respWrite http.ResponseWriter, request *http.Request) {
	// Functions serving as handlers take a http.ResponseWriter and a http.Request as arguments. The response writer is used to fill in the HTTP response.
	fmt.Fprintf(respWrite, "Hello!")
}

func main() {
	// Register handlers on server routes using the http.HandleFunc convenience function. It sets up the default router in the net/http package and takes a function as an argument.
	http.HandleFunc("/hello", sendHelloResponse)

	// Call the ListenAndServe with the port and a handler. nil tells it to use the default router.
	http.ListenAndServe(":8080", nil)
}