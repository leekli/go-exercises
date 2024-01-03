// Variadic functions
// Can be called with any number of trailing arguments
// Puts params into a slice[] so can be iterated

package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	// nums is now a slice[] of ints, so can iterate on it
	// using range..

	for _, num := range nums {
		total += num
	}

	fmt.Println("total is now:", total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// If the input is already a slice, invoke the fun with a trailing ...
	nums := []int{1, 2, 3, 4}
    sum(nums...)
}