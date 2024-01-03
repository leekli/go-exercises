// Recursion & recursive functions

package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1
    }

    return n * factorial(n-1)
}

func main() {
	// Factorial example
	fmt.Println("Factorial of 5 is:", factorial(5))

	// Fibonacci example with a closure involved
	var fib func(n int) int

    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n-1) + fib(n-2)
    }

    fmt.Println("Fibonacci number of 7 is:", fib(7))
}