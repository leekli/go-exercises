// Panic
// - Helping deal with unexpected errors. A panic typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.

package main

import "os"

func main() {
	// This immediately panics, and exits the program with a non-zero exit status - uncomment out to try
	// panic("A problem!")

	// Usually use panic to deal with errors
	// - Note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible.
	_, err := os.Create("/file")
    if err != nil {
        panic(err)
    }
}