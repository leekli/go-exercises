// Exit
// - Use os.Exit to immediately exit with a given status.

package main

import (
	"fmt"
	"os"
)

func main() {

	// defers will not be run when using os.Exit, so this fmt.Println will never be called.
    defer fmt.Println("!")

	// exit with status code 3
	// - If you’d like to exit with a non-zero status you should use os.Exit.
    os.Exit(3)
}