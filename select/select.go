// Select keyword
// Allows waiting on multiple channel operations, combining goroutines and channels

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create 2 channels
	c1 := make(chan string)
	c2 := make(chan string)

	// 2 anon funcs to run as goroutines async
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "msg one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "msg two"
	}()

	// Use select keyword with cases to deal with incoming messages from channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		}
	}
}