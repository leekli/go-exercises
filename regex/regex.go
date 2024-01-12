// Regular Expressions (Regex)
// - Go had built in Regex operations, there's a `regexp` package

package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Tests whether a pattern matches a string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("String match to peach?", match)

	// For other Regex tasks, need to use .Compile to create an optimised Regexp struct with methods on it
	regex, _ := regexp.Compile("p([a-z]+)ch")

	// Various methods in use below:
	fmt.Println(".MatchString:", regex.MatchString("peach"))

	// .FindString
	fmt.Println(".FindString:", regex.FindString("peach punch"))
	fmt.Println(".FindStringIndex:", regex.FindStringIndex("peach punch"))

	// .FindStringSubmatch
	// The Submatch variants include information about both the whole-pattern matches and the submatches within those matches. For example this will return information for both p([a-z]+)ch and ([a-z]+).
	fmt.Println(".FindStringSubmatch:", regex.FindStringSubmatch("peach punch"))
	fmt.Println(".FindStringSubmatchIndex:", regex.FindStringSubmatchIndex("peach punch"))

	// .FindAll
	fmt.Println(".FindAllString:", regex.FindAllString("peach punch pinch", -1))
	fmt.Println(".FindAllStringSubmatchIndex:", regex.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// Match Bytes
	fmt.Println("Match bytes:", regex.Match([]byte("peach")))

	// ReplaceAll
	fmt.Println("Replace peach with <fruit>:", regex.ReplaceAllString("a peach", "<fruit>"))
}