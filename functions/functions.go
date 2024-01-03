// Functions

package main

import "fmt"

// A function taking 2 int parameters, with a return value of an int in the definition/signature
func addNums(a int, b int) int {
	return a + b
}

// Function with multiple params of the same data type, can put together with 1 data type dec
func addMoreNums(a, b, c int) int {
	return a + b + c
}

func main() {
	// invoke `addNums`
	res1 := addNums(4, 5)
	fmt.Println("Result of addNums is:", res1)

	// invoke `addMoreNums`
	res2 := addMoreNums(10, 10, 11)
	fmt.Println("Result of addMoreNums is:", res2)
}