// WaitGroups
// - To wait for multiple goroutines to finish

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
	// This is a waitgroup put into 'wg' var
	var wg sync.WaitGroup

	// Launch 5 gorourtines, increment waitgroup by 1 each time so it knows how many waitgroups there are
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		// Using closure
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
	wg.Wait()
}