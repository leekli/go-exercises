// Closures & anonymous functions

package main

import "fmt"

func intSeq() func() int {
	i := 0
	
	// Return an anonymous function (creates the closure over var i)
	return func() int {
		i++

		return i
	}
}

func main() {
	// Create `nextInt` var which returns the first func from `intSeq`
	nextInt := intSeq()

	fmt.Println("1st invokation of nextInt() is:", nextInt())
	fmt.Println("2nd invokation of nextInt() is:", nextInt())
	fmt.Println("3rd invokation of nextInt() is:", nextInt())
	fmt.Println("4th invokation of nextInt() is:", nextInt())
	fmt.Println("5th invokation of nextInt() is:", nextInt())

	// Create new var invoking intSeq which will now be 2 lots of closure
	secondNextInt := intSeq()

	fmt.Println("1st invokation of secondNextInt() is:", secondNextInt())
}