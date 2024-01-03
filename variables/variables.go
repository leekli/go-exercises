// Variables

package main

import "fmt"

func main() {
	// 1 value using var (string)
	var a = "initial"
	fmt.Println(a)

	// Setting multiple values on one line (ints)
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Boolean
	var d = true
	fmt.Println(d)

	// Int
	var e int
	fmt.Println(e)

	// := shorthand syntax for declaring and initalising (can only be used inside func's)
	f := "apple"
	fmt.Println(f)
}