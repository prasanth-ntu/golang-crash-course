package main

import (
	"fmt"
	"math"
)

// Define interface //
type Shape interface {
	area() float64
}

// Let's use this interface with 2 structs //
type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Hola!")

	c1 := Circle{x: 0, y: 0, radius: 5}
	r1 := Rectangle{width: 10, height: 5}
	fmt.Println("c1 =", c1, "r1 =", r1)

	// Approach 1: Call area() method directly //
	fmt.Printf("Area of circle = %.2f\n", c1.area())
	fmt.Printf("Area of rectangle = %.2f\n", r1.area())

	// Approach 2: Call getArea() method //
	fmt.Printf("Area of circle = %.2f\n", getArea(c1))
	fmt.Printf("Area of rectangle = %.2f\n", getArea(r1))
}
