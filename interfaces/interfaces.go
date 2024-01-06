// Interfaces
// --> Named collections for method signatures

package main

import (
	"fmt"
	"math"
)

// Interface for 2 method signatures
type geometry interface {
	area() float64
	perim() float64
}

// 2 structs for a rectangle and circle
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}


// 4 functions for how to deal with area & perim depending if the struct is a rect or a circle
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Generic measure function which takes a rect OR circle, and invokes correct method depending on struct type
func measure(g geometry) {
	fmt.Println("given g:", g)
	fmt.Println("given g area:", g.area())
	fmt.Println("given g perim:", g.perim())
}

func main() {
	rect1 := rect{width: 10, height: 5}
	circ1 := circle{radius: 5}

	measure(rect1)
	measure(circ1)
}