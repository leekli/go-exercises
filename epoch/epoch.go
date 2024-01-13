// Epoch
// - Unix time: getting the number of seconds, milliseconds, or nanoseconds since the Unix epoch (January 1, 1970

package main

import (
	"fmt"
	"time"
)

func main() {
	// Get time right now
	now := time.Now()
	fmt.Println("Time now:", now)

	// Get time since Unix epoch
	fmt.Println("Time since unix epoch in seconds", now.Unix())
	fmt.Println("Time since unix epoch in milliseconds", now.UnixMilli())
	fmt.Println("Time since unix epoch in nanoseconds", now.UnixNano())

	// Convert epoch into corresponding time
	fmt.Println(time.Unix(now.Unix(), 0))
    fmt.Println(time.Unix(0, now.UnixNano()))
}