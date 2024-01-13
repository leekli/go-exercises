// Writing Files
// - Writing files in Go follows similar patterns to the ones we saw earlier for reading.

package main

import (
	"fmt"
	"os"
)

// Helper func for errors
func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
	// Dump a string (or just bytes) into a file.
    d1 := []byte("hello\ngo\n")
    err := os.WriteFile("./file2.txt", d1, 0644)
    check(err)

	//F or more granular writes, open a file for writing.
    f, err := os.Create("/tmp/dat2")
    check(err)

	// It’s idiomatic to defer a Close immediately after opening a file.
    defer f.Close()

	// Can Write byte slices
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

	// Can use WriteString to write an actual string
    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n3)

	// Issue a Sync to flush writes to stable storage.
    f.Sync()
}