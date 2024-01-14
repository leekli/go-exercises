// HTTP Client
// - The Go standard library comes with support for HTTP clients and servers in the net/http package.

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// Initalise a HTTP GET Request & error handle
	response, err := http.Get("https://example.org/")

	if err != nil {
		panic(err)
	}

	// End the response when enclosing main func ends
	defer response.Body.Close()

	// Print status code responded with
	fmt.Println("Response status:", response.Status)

	// To print response body
	scanner := bufio.NewScanner(response.Body)

    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}