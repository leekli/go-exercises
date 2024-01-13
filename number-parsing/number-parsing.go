// Number Parsing
// - Parsing numbers from strings
// - using `strconv` package

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Parse a float
    floatParse, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(floatParse)

	// Parse int
    intParse, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(intParse)

	// Parse hex
    hexParse, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(hexParse)

	// Parse an unsigned int (positive number)
    uIntParse, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(uIntParse)

	// Atoi is base-10 int parsing
    base10Parse, _ := strconv.Atoi("135")
    fmt.Println(base10Parse)

	// What happens when pass a string to Atoi for base-10 parsing
    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}