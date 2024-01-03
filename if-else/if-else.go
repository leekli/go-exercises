// If, else if, else statements

package main

import "fmt"

func main() {
	// If, else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is od")
	}

	// Just an if
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Using logical operators (|| and &&)
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// Initalising a variable in if statement, which is then available within the block
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}