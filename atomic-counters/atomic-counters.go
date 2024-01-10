// Atomic Counters
// - uses 'sync/atomic' package
// - The primary mechanism for managing state in Go is communication over channels. We saw this for example with worker pools. There are a few other options for managing state though. Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// Create an atomic counter of Uint64 type (unsigned positive only numbers up to 64-bit)
	var counter atomic.Uint64

	// Create a WaitGroup to wait for all goroutines to finish
	var waitgroup sync.WaitGroup

	// For loop to start up 50 goroutines, each one increasing the atomic counter 'couner' by 1000 (due to inner loop), also increases waitgroup by 1 for each new gochannel created so that waitgroup can track all goroutines in operation
	for i := 0; i < 50; i++ {
		// Add 1 to waitgroup
		waitgroup.Add(1)

		// anon func to start a goroutine
		go func() {
			for c := 0; c < 1000; c++ {
				counter.Add(1)
			}

			waitgroup.Done()
		}()
	}

	// waitgroup needs to wait for all outstanding waitgroups/goroutines to finish (will be 50 in total here)
	waitgroup.Wait()

	fmt.Println("Counter is:", counter.Load())
}