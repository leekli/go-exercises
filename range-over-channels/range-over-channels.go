// Range over Channels
// - using `range` iteration to iterate over values received from a channel instead of `for` keyword

package main

import "fmt"

func main() {

	// Create channel which will take 2 values
    queue := make(chan string, 2)

	// Send the two values to channel
    queue <- "one"
    queue <- "two"

	// Close the channel
    close(queue)

	// Use `range` to iterate over values receives from the `queue` channel above
    for element := range queue {
        fmt.Println("Received:", element)
    }
	// - Note: because used `close` method above, the range loop knows to stop iterating, if there wasn't a close then an error would be produced
}