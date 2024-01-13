// Line Filters
// - To take in input from stdio, process it and print out, use of stdout and stderr

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

    scanner := bufio.NewScanner(os.Stdin)

	// Wrapping the unbuffered os.Stdin with a buffered scanner gives us a convenient Scan method that advances the scanner to the next token; which is the next line in the default scanner.
    for scanner.Scan() {
		
        ucl := strings.ToUpper(scanner.Text())

        fmt.Println(ucl)
    }

	// Check for errors during Scan. End of file is expected and not reported by Scan as an error.
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}