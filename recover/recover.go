// Recover
// - Ability to recover from a `panic` command
// - Go makes it possible to recover from a panic, by using the recover built-in function. A recover can stop a panic from aborting the program and let it continue with execution instead.
// - An example of where this can be useful: a server wouldn’t want to crash if one of the client connections exhibits a critical error. Instead, the server would want to close that connection and continue serving other clients. In fact, this is what Go’s net/http does by default for HTTP servers.

package main

import "fmt"

func mayPanic() {
	panic("Problemo!")
}

func main() {
	// need to defer the recover function so it executes at the end of the enclosing function; in this case 'main'
	defer func() {
		if recover_err := recover(); recover_err != nil {
			fmt.Println("Recovered. Error:\n", recover_err)
		}
	}()

	mayPanic()

	fmt.Println("This line executes after the recovery showing it worked")
}