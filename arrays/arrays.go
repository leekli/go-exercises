// Arrays

package main

import "fmt"

func main() {
	// Empty array of size 5 (using var)
	var a [5]int
	fmt.Println("empty array:", a)

	// Add an element to 1 of the above array
	a[4] = 100
	fmt.Println("array now:", a)
	fmt.Println("element 4:", a[4])

	// Length of array
	fmt.Println("Array length", len(a))

	// Create array and fill all indexes (using :=)
	b := [5]int{1,2,3, 4,5}
	fmt.Println("b array:", b, "with length:", len(b))

	// Create a 2d array with 2 for loops
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d array:", twoD, "with length:", len(twoD))
}