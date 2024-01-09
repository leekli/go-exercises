// Timers
// - We often want to execute Go code at some point in the future, or repeatedly at some interval. Go’s built-in timer and ticker features make both of these tasks easy.
// - They form part of Go's concurrency mechanism (Goroutines and channels), not inheriently asynchronous

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a timer for 2 seconds
	timer1 := time.NewTimer(1 * time.Second)

	// Blocks timers channel until a value is sent
	<-timer1.C
	fmt.Println("Timer 1 fired!")

	// Can actually stop timers...
	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 was stopped!")
	}

	// Can use sleep to pause (in this example will pause the program from ending for 2 seconds)
	time.Sleep(2 * time.Second)
}