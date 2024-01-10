// Rate Limiting
// - Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service. Go elegantly supports rate limiting with goroutines, channels, and tickers.

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a requests channel, 5 messages max
	requests := make(chan int, 5)

	// Send the 5 messages to the requests channel then invoke close to say we've finished sending all requests
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Create a time channel, will receive a value every 1 second
	limiter := time.Tick(1000 * time.Millisecond)

	// Now run through all the request channel messages and print them, using limiter of 1 message per 1 second
	for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }
}