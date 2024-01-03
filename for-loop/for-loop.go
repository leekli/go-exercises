// For loops

package main

import "fmt"

func main() {
	i := 1

	// With just a condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// With all 3 steps of initialisation, stopping condition and increment step
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Without steps, continues until a break
	for {
		fmt.Println("loop")
		break
	}

	// Using continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
  }