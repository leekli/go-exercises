// Pointers (Passing references to values)

package main

import "fmt"

// Function creates a copy of arg given
func zeroval(ival int) {
	ival = 0
}

// Function has a pointer to the argument given, dereferences memory address on line 14
func zeroptr(iptr *int) {
	*iptr = 0
}

func main () {
	i := 1
	fmt.Println("Inital value of main i is:", i)

	// Pass i to non-pointer `zeroval` func, i gets a copy created in the func, main i not affected (passed by value, not ref)
	zeroval(i)
	fmt.Println("Value of main i after zeroval invoked:", i)

	// Pass i as a pointer to `zeroptr` func, main i is affected as value will be changed directly by ref
	// & syntax gives memory address pointer
	zeroptr(&i)
	fmt.Println("zeroptr return value/i is now:", i)

	// Print memory address of i using %
	fmt.Println("Memory address of main i is:", &i)
}