// Logging
// - The Go standard library provides straightforward tools for outputting logs from Go programs, with the log package for free-form output and the log/slog package for structured output.
// - Using `log` package

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	// Similiar functions from Go are available in the `log` package which format output appropiate for logging
	log.Println("Standard Log:")

	// Can use flags for setting format of time output:
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
    log.Println("With microseconds included")

	// Flag for adding the file name and line number where log came from:
	log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("With file name & line number included")

	// To create a customer logger and pass it around:
	mylog := log.New(os.Stdout, "Lee's Log: ", log.LstdFlags)
    mylog.Println("from mylog")

	// To change the Prefit
	mylog.SetPrefix("Lee's New Logger: ")
    mylog.Println("from mylog")

	// Loggers can have custom output targets; any io.Writer works.
	var buf bytes.Buffer
    buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("hello")

	fmt.Print("from buflog: ", buf.String())
}