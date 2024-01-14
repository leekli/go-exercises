// Command Line Flags
// - are a common way to specify options for command-line programs. For example, in wc -l the -l is a command-line flag.
// - available using `flag` package in Go for command line parsing

package main

import (
	"flag"
	"fmt"
)

func main() {
	// Set up 3 flag declarations (a string, int and boolean types) with default values, returns a POINTER
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")

	// It’s also possible to declare an option that uses an existing var declared elsewhere in the program. Note that we need to pass in a pointer to the flag declaration function.
	var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	// Print values
	// - Note that we need to dereference the pointers with e.g. *wordPtr to get the actual option values.
	fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}