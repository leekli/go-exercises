// Go Channel Synchronization
// - sync execution across goroutines

package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working... ")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {
	// Create channel
	done := make(chan bool, 1)
	
	// Invoke func with goroutine concurrently.
	go worker(done)

	// Block & Wait until the channel receives a notification (from line 16)
	// note: don't need put into a msg var, could just be <- done
	msg := <- done
	fmt.Println("Output of done channel: ", msg)
}