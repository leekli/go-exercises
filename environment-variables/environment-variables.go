// Environment Variables
// - Environment variables are a universal mechanism for conveying configuration information to Unix programs.
// - using `os` package

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// To set a key-value pair env var
	os.Setenv("FOO", "1")
	fmt.Println("Env var - FOO:", os.Getenv("FOO"))
	fmt.Println("[env var doesnt exist] BAR:", os.Getenv("BAR"))

	// Get a system env var which already exists
	fmt.Println("Env var - $HOME:", os.Getenv("HOME"))

	// Use os.Environ to list all key/value pairs in the environment. 
	fmt.Println()
    for _, envVar := range os.Environ() {
        pair := strings.SplitN(envVar, "=", 2)
        fmt.Println(pair[0])
    }
}