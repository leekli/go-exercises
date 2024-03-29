// Contexts
// - HTTP servers are useful for demonstrating the usage of context.Context for controlling cancellation. A Context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines.

package main

import (
	"fmt"
	"net/http"
	"time"
)

func sendHelloResponse(w http.ResponseWriter, req *http.Request) {
	// A context.Context is created for each request by the net/http machinery, and is available with the Context() method.
    ctx := req.Context()
	
    fmt.Println("Server: hello handler started")
    defer fmt.Println("Server: hello handler ended")

	// Wait for a few seconds before sending a reply to the client. This could simulate some work the server is doing. While working, keep an eye on the context’s Done() channel for a signal that we should cancel the work and return as soon as possible. (CHANNELS)
    select {
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done():
		// The context’s Err() method returns an error that explains why the Done() channel was closed.
        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }
}

func main() {

    http.HandleFunc("/hello", sendHelloResponse)
    http.ListenAndServe(":8081", nil)
}