// Functions: Returning multiple values
// Function signature needs return values in brackets to say number of returned values e.g. (x, x)

package main

import "fmt"

// Return 2 values, needs func signature, and return values on line 8 are comma-seperated
func vals() (int, int) {
	return 3, 7
}

func main() {
	// Access multiple return values by declaring 2 variables comma-seperated
	a, b := vals()
	fmt.Println("return value a is:", a)
	fmt.Println("return value b is:", b)

	// If we only want the 2nd value returned, ignore first with a _ (blank identifier)
	_, c := vals()
	fmt.Println("return value c is:", c)
}