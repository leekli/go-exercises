// Struct Embedding
// Can embed structures and interfaces to support composition of types

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	container1 := container{base: base{num: 2}, str: "Hello!"}

	fmt.Printf("container1={num: %v, str: %v}\n", container1.num, container1.str)

	fmt.Println("can also access num directly:", container1.base.num)

	fmt.Println("describe:", container1.describe())

	type describer interface {
        describe() string
    }

    var d describer = container1
    fmt.Println("describer:", d.describe())
}