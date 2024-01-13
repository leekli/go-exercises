// JSON
// - Built in support for JSON, package called `encoding/json`

package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page 	int			`json:"page"`
	Fruits	[]string	`json:"fruits"`
}

func main() {
	// Using JSON on basic data types: bool, int, float, string:
	boolA, _ := json.Marshal(true)
	fmt.Println("Boolean:", string(boolA))

	intA, _ := json.Marshal(2)
	fmt.Println("Int:", string(intA))

	floatA, _ := json.Marshal(2.34)
	fmt.Println("Float:", string(floatA))

	stringA, _ := json.Marshal("Hello!")
	fmt.Println("String:", string(stringA))

	// Using JSON on slices and maps:
	fruitsSlice := []string{"apple", "peach", "pear"}
	sliceA, _ := json.Marshal(fruitsSlice)
	fmt.Println("Fruits slice:", string(sliceA))

	fruitsMap := map[string]int{"apple": 1, "banana": 2, "orange": 3}
	mapA, _ := json.Marshal(fruitsMap)
	fmt.Println("Fruits Map:", string(mapA))

	// Using JSON with custom data type but without the `json:` exports
	res1 := &response1{
		Page: 1,
		Fruits: []string{"apple", "banana", "orange"},
	}
	res2, _ := json.Marshal(res1)
	fmt.Println("Custom data type without export:", string(res2))

	// Using JSON with custom data type and `json` export keys
	res3 := &response2{Page: 2, Fruits: []string{"orange", "apple", "strawberry"}}
	res4, _ := json.Marshal(res3)
	fmt.Println("Custom data type with json export:", string(res4))

	// Decoding JSON (convert bytes to map):
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	json.Unmarshal(byt, &dat)
	fmt.Println(dat)
	num := dat["num"].(float64)
    fmt.Println(num)
}