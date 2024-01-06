// Structs
// Typed collection of fields, grouping data together, mix of data types (OOP equiv)

package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

func main() {
	// Initalise struct with 2 values
	fmt.Println(person{"Bob", 20})

	// Initalise struct with 2 values with key names explicit
	fmt.Println(person{name: "Alice", age: 30})

	// Initalise struct with only 1 of 2 values (age becomes zero indexed)
	fmt.Println(person{name: "Lee"})

	// Using a pointer
	fmt.Println(&person{name: "Ann", age: 30})

	// use constructor function to initalise struct
	fmt.Println(newPerson("Lee"))

	// Access keys with dot notation
	person1 := newPerson("Lee")
	fmt.Println("name key", person1.name)
	fmt.Println("age key", person1.age)

	// Update a key with dot notation
	person1.age = 35
	fmt.Println("updated age key", person1.age)
}