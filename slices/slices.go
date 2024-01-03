// Slices (More powerful, flexible and in use than Arrays), collection of same data type in a list
// Set a slice up similiar to an array but don't set a length

package main

import (
	"fmt"
	"slices"
)

func main() {
	// Empty slice
	var s []string
	fmt.Println("slice s:", s, "is s nil?...", s == nil, "s length is 0??:", len(s) == 0)

	// Use make() to create an empty slice with non-zero length (sets a capacity)
	s = make([]string, 3)
	fmt.Println("empty slice of 3:", s, "s length:", len(s), "capacity:", cap(s))

	// Set and get slice values just like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s contents:", s)
	fmt.Println("s[0]:", s[0])
	fmt.Println("s[1]:", s[1])
	fmt.Println("s[2]:", s[2])
	fmt.Println("s length:", len(s))

	// Use append() method to add a new element, creates copy of slice, needs a return value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("new s:", s, "new len:", len(s), "new cap:", cap(s))

	// Copy slice using copy(dest, src) method (use size/len of slice above to make a copy)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied slice:", c, "copied slice len:", len(c), "copied slice cap:", cap(c))

	// Using slice operator, creates a copy of new slice by index
	l := s[2:5]
    fmt.Println("slice from index 2-5 excl. 5:", l)

	l = s[:5]
    fmt.Println("slice up to but excl 5 :", l)

	l = s[2:]
    fmt.Println("slice from index and including 2:", l)

	// Declaring AND initalisting a slice
	t := []string{"hey", "hello", "hi"}
	fmt.Println("slice t with 3 strings:", t, "t len:", len(t), "t cap:", cap(t))

	// Using the 'slice' package gives extra operations as checking if 2 slices are equal
	t2 := []string{"hey", "hello", "hi"}
	if slices.Equal(t, t2) {
		fmt.Println("t is equal to t2")
	}

	// 2d slices composition with varying lengths
	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}