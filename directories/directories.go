// Directories
// - Go has several useful functions for working with directories in the file system.

package main

import (
	"fmt"
	"os"
)

// Helper error handling function
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// To create a new directory, give it's name and permissions number
	err := os.Mkdir("subdir", 755)
	check(err)

	// When creating temporary directories, it’s good practice to defer their removal. os.RemoveAll will delete a whole directory tree (similarly to rm -rf). - will run at the end of the enclosing function
	defer os.RemoveAll("subdir")

	// Create a file within the `subdir` directory, can set permissions with the WriteFile
	createEmptyFile := func(name string) {
        d := []byte("")
        check(os.WriteFile(name, d, 0644))
    }

    createEmptyFile("subdir/file1")

	// We can create a hierarchy of directories, including parents with MkdirAll. This is similar to the command-line mkdir -p.
	err = os.MkdirAll("subdir/parent/child", 0755)
    check(err)

	createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

	// ReadDir lists directory contents, returning a slice of os.DirEntry objects.
	c, err := os.ReadDir("subdir/parent")
    check(err)

    fmt.Println("Listing subdir/parent")

    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

	// Chdir lets us change the current working directory, similarly to cd.
	// - Now we’ll see the contents of subdir/parent/child when listing the current directory.
	err = os.Chdir("subdir/parent/child")
    check(err)

	c, err = os.ReadDir(".")
    check(err)

	fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

	err = os.Chdir("../../..")
    check(err)
}