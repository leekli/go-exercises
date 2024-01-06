// Strings & Runes

// In Go; strings are read-only slice of bytes, text encoded in utf-8 (not by character).
// Concept of a character is called a 'rune' - number representing Unicode code point.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const helloStr = "Hello!"

	fmt.Println("Lenngth of helloStr", len(helloStr))

	// For loop produces hex value for each byte in string
	for i := 0; i < len(helloStr); i++ {
		fmt.Printf("%x ", helloStr[i])
	}
	fmt.Println()

	// Rune count (each character count)
	fmt.Println("Rune count:", utf8.RuneCountInString(helloStr))

	// Range loop with string and runes
	for idx, runeValue := range helloStr {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
 }