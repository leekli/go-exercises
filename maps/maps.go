// Maps
// Associative data type (hash/dict/obj in other languages), key-value pairs

package main

import (
	"fmt"
	"maps"
)

func main() {
	// Use make() function to create a Map
	m := make(map[string]int)

	// Assign key/value pairs
	m["key1"] = 2
	m["key2"] = 5

	fmt.Println("m Map:", m)

	// Get a value from a key name
	fmt.Println("m['key1'] is:", m["key1"])

	v2 := m["key2"]
	fmt.Println("m['key2'] is:", v2)

	// Use len() to get the number of key/value pairs in a map
	fmt.Println("len of map is:", len(m))

	// Use delete() to delete a key/value pair from a map
	delete(m, "key1")
	fmt.Println("map is now:", m, "with new len of after deletion:", len(m))

	// To remove all key/value pairs and empty the map, use clear()
	clear(m)
	fmt.Println("map is now:", m, "with new len of after deletion:", len(m))

	// Optional second return value when getting a value from map to work out if missing key or is 0/""
	_, prs := m["key2"]
    fmt.Println("prs:", prs)

	// Using := in a func to declare and initalise a map
	map2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map2 is:", map2, "with len:", len(map2))

	// Use 'maps' package for more functionality/methods
	map3 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(map2, map3) {
        fmt.Println("map2 is equal to map3")
    }
}