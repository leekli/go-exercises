// Mutexes
// - Like atomic counters for managing simple state.
// - For more complex state we can use a mutex to safely access data across multiple goroutines.
// Use lock and unlock mechanisms to allow/deny access to data to ensure multiple sources don't try change data at same time

package main

import (
	"fmt"
	"sync"
)

// Create a struct for containing the data which multiple gochannels may want access, need to add a mutex property to signify this is a mutex
// MUST only use pointers to mutexes, cannot be copied
type Container struct {
	mu sync.Mutex
	counters map[string]int
}

// This function will lock the container/data, apply the change to the data, then unlock it - prevening other gochannels from trying to access the same data at the same time. Uses received parameter to bind to the 'Container' struct with a pointer
func (container *Container) increment(name string) {
	// Lock the container
	container.mu.Lock()

	// Unlock the container but apply 'defer' keyword prefix
	// defer... Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
	defer container.mu.Unlock()

	// Increase container map key-value by 1
	container.counters[name]++
}

func main() {
	// Create instance of struct
	container := Container {
		counters: map[string]int{"a": 0, "b": 0},
	}

	// Set up a WaitGroup for the multiple goroutines
	var waitgroup sync.WaitGroup

	// Function for increasing the counter n number of times
	doIncrement := func(name string, num int) {
		for i := 0; i < num; i++ {
			container.increment(name)
		}

		// Send signal to WaitGroup that this is done
		waitgroup.Done()
	}

	// Now add 3 routines to the waitgroup, and send 3 goroutines off to increment the Map counter data
	waitgroup.Add(3)
	go doIncrement("a", 1000)
	go doIncrement("a", 1000)
	go doIncrement("b", 1000)

	// Make WaitGroup wait for all async goroutines to finish
	waitgroup.Wait()

	// Print results
	fmt.Println("Total counters now are:", container.counters)
}