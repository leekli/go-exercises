// Non-Blocking Channel Operations
// - Basic sends/receives on channels are blocking, can use select and default keywords to non-block sends, receives, etc.

package main

import (
	"fmt"
)

func main() {
	// Create 2 channels
	messages := make(chan string)
	signals := make(chan bool)

	// Select to check if any messages RECEIVED
	select {
	case msg := <-messages:
		fmt.Println("Message received:", msg)
	default:
		fmt.Println("No messages received")
	}

	// Select to check if any messages SENT
	msg := "Hello!"
	select {
	case messages <- msg:
		fmt.Println("Sent message:", msg)
	default:
		fmt.Println("No messages sent")
	}

	// Use select on both channels (multi-way non-blocking select)
	select {
	case msg := <-messages:
        fmt.Println("Message received:", msg)
    case sig := <-signals:
        fmt.Println("Signal received:", sig)
    default:
        fmt.Println("No activity!")
	}
}