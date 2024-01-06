// Methods on struct types

package main

import "fmt"

type rect struct {
	width, height int
}

// Using method receiver types
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("r area:", r.area())
	fmt.Println("r perim:", r.perim())

	// Using pointer
	r2 := &r
	fmt.Println("r2 area:", r2.area())
	fmt.Println("r2 perim:", r2.perim())
}