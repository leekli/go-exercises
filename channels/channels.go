// Channels
// - Pipes to connect concurrent goroutines. Send values from one goroutine and receive in another.

package main

import "fmt"

func main() {
	// Creates a channel
	messages := make(chan string)

	// Anon func with goroutine
	go func() { messages <- "ping" }()

	// Receives the messages and prints
	msg := <- messages
	fmt.Println(msg)
}