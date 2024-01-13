// URL Parsing
// - `net/url` package includes parse utility .Parse

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// Generate a standard url string
	urlStr := "postgres://user:pass@host.com:5432/path?k=v#f"

	// using .Parse method 
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Full parsed url:", parsedUrl)
	fmt.Println("Scheme from url:", parsedUrl.Scheme)
	fmt.Println("User property from url:", parsedUrl.User)
	fmt.Println("Host from url:", parsedUrl.Host)

	host, port, _ := net.SplitHostPort(parsedUrl.Host)
    fmt.Println("Just the host:", host)
    fmt.Println("Just the port:", port)

	fmt.Println("Path from url:", parsedUrl.Path)
	fmt.Println(parsedUrl.Fragment)
}