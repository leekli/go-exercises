// File Paths
// - The `filepath` package provides functions to parse and construct file paths in a way that is portable between operating systems; dir/file on Linux vs. dir\file on Windows, for example.

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// To create a new filepath in a quick way use Join if there's multiple sub-paths
	// - ALWAYS use .Join instead of concatenating them
	// - Join will also normalize paths by removing superfluous separators and directory changes.
	path := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("Joined Filepath:", path)

	// To get the path (excluding filename)
	fmt.Println("Dir(p):", filepath.Dir(path))

	// To get the file base
    fmt.Println("Base(p):", filepath.Base(path))

	// Can find out if a filepath is absolute or relative (returns boolean)
	fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))

	fileName := "config.json"
	// To split out extension from file:
	extension := filepath.Ext(fileName)
    fmt.Println(extension)

	// And... to get the file name suffix
	fmt.Println(strings.TrimSuffix(fileName, extension))
}