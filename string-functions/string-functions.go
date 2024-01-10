// String Functions in Go
// - `strings` package available in standard library which has a number of string functions for operations on strings

package main

// this imports the string package as the variable alias 's
import (
	"fmt"
	s "strings"
)

// this aliases fmt.Println to `p` for a short hand, point to function ref (not invoked)
var  p = fmt.Println

func main() {
	// Various string functions available 👇
	p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
}