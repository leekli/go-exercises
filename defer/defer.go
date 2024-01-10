// Defer
// - Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.

package main

import (
	"fmt"
	"os"
)

func main() {
	// 3 functions to create a file, write to that file then close that file. We can defer closing the file, and should do Immediately after getting a file object with createFile, we defer the closing of that file with closeFile. This will be executed at the end of the enclosing function (main), after writeFile has finished.

	file := createFile("./defer.txt")
	defer closeFile(file)
	writeFile(file)
}

func createFile(path string) *os.File {
	fmt.Println("Creating your file...")

	newFile, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	fmt.Println("File created.")

	return newFile
}

func writeFile(file *os.File) {
	fmt.Println("Writing to your file...")

	fmt.Fprintf(file, "I am data")

	fmt.Println("File write complete.")
}

func closeFile(file *os.File) {
	fmt.Println("Closing your file...")

	err := file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
	}

	fmt.Println("File closed.")
}