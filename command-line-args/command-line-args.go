// Command Line Arguments
// - Command-line arguments are a common way to parameterize execution of programs. For example, `go run hello.go` uses run and hello.go arguments to the go program.
// - argument parsing available from `os` package

package main

import (
	"fmt"
	"os"
)

func main() {
	// Entire absolute path with args attached
	argsWithProg := os.Args

	// This grabs all arguments after the filename
	argsWithoutProg := os.Args[1:]

	// Get individual args with normal indexing
	arg := os.Args[3]

	fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}