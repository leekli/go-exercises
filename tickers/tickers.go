// Tickers (part of time)
// - Timers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals.
// - Tickers use a similar mechanism to timers: a channel that is sent values.

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a new Ticker (a ticker for every 500 milliseconds)
	ticker := time.NewTicker(500 * time.Millisecond)

	// Create a channel to track when all values have been sent (no more to send) to channel
	done := make(chan bool)

	// Use goroutine and select to await the values as they arrive every 500ms
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick received at", t)
			}
		}
	}()

	// Need a delay set to make sure anon func above has time to get the incoming values 
	time.Sleep(1600 * time.Millisecond)

	// Stop the ticker and send value to done channel to say no more values will be sent
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}