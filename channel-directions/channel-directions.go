// Channel directions
// Func params can be set up to say whether they are meant to only send OR receive values via channels, increasing type safety

package main

import "fmt"

// Ping func only accepts a channel for sending values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Pong func accepts 1 channel for receives (ping) values from a channel, and a second to send (pongs) values to a channel
func pong(pings <-chan string, pongs chan<- string) {
	// Listens for messages from pings channel
	msg := <-pings

	// Then sends pong to pongs channel once ping received
	pongs <- msg
	pongs <- "pong"
}

func main() {
	// Set up 2 channels
	pings := make(chan string, 1)
	pongs := make(chan string, 2)

	ping(pings, "ping")
	pong(pings, pongs)

	fmt.Println(<-pongs)
	fmt.Println(<-pongs)
}