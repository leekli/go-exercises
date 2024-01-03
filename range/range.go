// Ranges
// `range` iterates over elements in a variety of data structures.

package main

import "fmt"

func main() {
	// Use range on a slice of numbers to add them all up (use _ to ignore first return value which is the index)
	nums := []int{2,3,4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum is now:", sum)

	// Use range and get the index also
	for idx, num := range nums {
		if num == 3 {
			fmt.Println("index of 3 is:", idx)
		}
	}

	// Use range on a map to get key/value pairs using printf and formatting verbs
	kvs := map[string]string{"a": "apple", "b": "banana", "c": "carrot"}

	for k, v := range kvs {
		fmt.Printf("%s --> %s\n", k, v)
	}

	// Use range just to get the keys
	for k := range kvs {
		fmt.Println("Key is: ", k)
	}

	// Use range just to get the values only (ignore keys with _)
	for _, v := range kvs {
		fmt.Println("Value is:", v)
	}

	// Using range on a string returns index and the rune (ascii code)
	s := "hello"
	for idx, char := range s {
		fmt.Println("char code is:", char, "at index:", idx)
	}
}