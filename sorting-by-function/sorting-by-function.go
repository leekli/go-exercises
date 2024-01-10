// Sorting by Funcions
// - Sometimes we’ll want to sort a collection by something other than its natural order. For example, suppose we wanted to sort strings by their length instead of alphabetically.

package main

import (
	"cmp"
	"fmt"
	"slices"
)


func main() {
	// Slice of strings of varying length
	fruits := []string{"jsdfhaasfssj", "banana","apple", "kiwi"}

	// Declare a custom function to do the comparison between two elements
	lengthCompare := func(word1, word2 string) int {
		return cmp.Compare(len(word1), len(word2))
	}

	// Apply this using .SortFunc from slices package - pass the slice and function reference & print result
	slices.SortFunc(fruits, lengthCompare)

	fmt.Println("Fruits slice is now (after custom sorting):", fruits)

	// Can also do this on things that aren't built-in types like strings or numbers, use structs or specific keys to compare against
	type Person struct {
		name string
		age int
	}

	people := []Person{
		{name: "Jax", age: 37},
        {name: "TJ", age: 25},
        {name: "Alex", age: 72},
	}

	compareAges := func(personA, personB Person) int {
		return cmp.Compare(personA.age, personB.age)
	}

	slices.SortFunc(people, compareAges)

	fmt.Println("People slice is now (after custom sorting):", people)
}