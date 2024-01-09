// Timeouts
// - Important for programs that connect to external resources or that otherwise need to bound execution time.
// - Uses channels and `select`

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create channel with 1 value arg & invoke with anon func goroutine
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Use select keyword to deal with first receive that's ready or timeout
	select {
	case result := <-c1:
		fmt.Println("result:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out after 1 second")
	}

	// Do the above again with a 2nd channel, send value back within time
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case result := <-c2:
		fmt.Println("result:", result)
	case <-time.After(3 * time.Second):
		fmt.Println("Timed out after 1 second")
	}
}