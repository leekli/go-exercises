// Channel Buffering
// - Allows limited number of values without a corresponding receiver for those values

package main

import "fmt"

func main() {
	// Set up a channel, 2nd argument is number of values it can take
	messages := make(chan string, 2)

	messages <- "ping"
	messages <- "pong"

	fmt.Println("Message 1", <- messages)
	fmt.Println("Message 2", <- messages)
}