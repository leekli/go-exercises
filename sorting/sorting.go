// Sorting
// - Go 'slices' package can help sort data

package main

import (
	"fmt"
	"slices"
)

func main() {
	// Create slice of strings
	strs := []string{"d", "a", "b", "c", "e"}

	// Sort the slice of strings and print result
	slices.Sort(strs)
	fmt.Println("strs slice is now (after sorting):", strs)

	// Slice of numbers, then sort them and print
	nums := []int{6, 3, 1, 2, 5, 4}
	slices.Sort(nums)
	fmt.Println("nums slice is now (after sorting):", nums)

	// Use .IsSorted to return a boolean of if a slice is sorted or not
	areStrsSorted := slices.IsSorted(strs)
	fmt.Println("Is strs slice sorted?", areStrsSorted)

	areNumsSorted := slices.IsSorted(nums)
	fmt.Println("Is nums slice sorted?", areNumsSorted)

	moreNums := []int{6, 3, 1, 2, 5, 4}
	areMoreNumsSorted := slices.IsSorted(moreNums)
	fmt.Println("Is moreNums slice sorted?", areMoreNumsSorted)
}