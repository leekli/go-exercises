// Reading Files
// - Reading and writing files are basic tasks needed for many Go programs.

package main

import (
	"fmt"
	"os"
)

// Reading files requires error handling, helper func to use for it:
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Read from a file, 2 return values: the read file contents or an error (opens the file, puts contents in memory, closes file)
	data, err := os.ReadFile("./test.txt")
    check(err)
    fmt.Print("Output: ", string(data))

	// Opening a file
	file, err := os.Open("./test.txt")
	check(err)
	
	// Read the first 5 bytes
	b1 := make([]byte, 5)
    n1, err := file.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Can also Seek to a known location in the file and Read from there.
	o2, err := file.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := file.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: ", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))

	// To close the file
	file.Close()
}