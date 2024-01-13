// Random Numbers
// - using `math/rand` package

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 2 Random numbers 0-100 seperated by a new line:
	fmt.Print("Random number 1: ", rand.Intn(100), ", ")
	fmt.Print("Random number 2: ", rand.Intn(100))
    fmt.Println()

	// Generate a random float
	fmt.Println("Random float:", rand.Float64())

	// 2 random floats
	fmt.Print("Random float 1: ", (rand.Float64()*5)+5, ", ")
    fmt.Print("Random float 2: ", (rand.Float64() * 5) + 5)
    fmt.Println()
}