// Temporary Files & Directories
// - Throughout program execution, we often want to create data that isn’t needed after the program exits. Temporary files and directories are useful for this purpose since they don’t pollute the file system over time.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Error helper func
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// The easiest way to create a temporary file is by calling os.CreateTemp. It creates a file and opens it for reading and writing. We provide "" as the first argument, so os.CreateTemp will create the file in the default location for our OS.
	tempFile, err := os.CreateTemp("", "sample")
	check(err)

	fmt.Println("TempFile file name:", tempFile.Name())

	// This does the cleanup of the temp file @ the end of execution of main()
	defer os.Remove(tempFile.Name())

	// Can create a temp dir to write many files to:
	// - Write some data:
	_, err = tempFile.Write([]byte{1, 2, 3, 4})
    check(err)

	// - Create temp dir
	dname, err := os.MkdirTemp("", "sampledir")
    check(err)
    fmt.Println("Temp dir name:", dname)

	// - Clean up at the end of enclosing finc
	defer os.RemoveAll(dname)

	// - Now can synthesize temporary file names by prefixing them with our temporary directory.
	fname := filepath.Join(dname, "file1")
    err = os.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)
}