// Time

package main

import (
	"fmt"
	"time"
)

func main() {
	// Shorthand fmt.Println
	p := fmt.Println

	// Get current time
	now := time.Now()
	p("Time now:", now)

	// Building a time struct by passing year, month, day, etc.
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

	// Extrac from the `then` struct:
	p("Year:", then.Year())
    p("Month:", then.Month())
    p("Day:", then.Day())
    p("Hour:", then.Hour())
    p("Minute:", then.Minute())
    p("Second:", then.Second())
    p("Nanosecond:", then.Nanosecond())
    p("Location:", then.Location())

	// Get the weekday name
	p("Weekday name:", then.Weekday())

	// These methods compare two times, testing if the first occurs before, after, or at the same time as the second, respectively.
	p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

	// Getting duration between two intervals along with various methods for showing the difference
	diff := now.Sub(then)
    p(diff)
	p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())
}