// Goroutines
// - A lightweight thread of execution

package main

import (
	"fmt"
	"time"
)

func testFunc(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Call testFunc directly sync
	testFunc("direct")

	// Call testFunc concurrently with goroutine
	go testFunc("goroutine")

	// Go routine for anonymous func
	go func(msg string) {
        fmt.Println(msg)
    }("going")

	// Need this to wait for them to finish
	time.Sleep(time.Second)
    fmt.Println("done")
}