package main

import (
	"fmt"
	"math"
)

type Shapes interface {
	Perimeter() float64
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
        length float64
	width float64
}

func (r Rectangle) Perimeter() float64 {
        return 2 * (r.length + r.width)
}

func (r Rectangle) Area() float64 {
        return r.length * r.width
}

type Square struct {
        length float64
}

func (s Square) Perimeter() float64 {
        return 4 * s.length
}

func (s Square) Area() float64 {
        return s.length * s.length
}

func CalculateShapeCoverage(shape Shapes) {
	switch s := shape.(type) {
	case 	Circle, 
		Rectangle, 
		Square:
		fmt.Printf("Perimeter: %.2f (m), Area: %.2f (m^2)\n",
					s.Perimeter(),  s.Area())
	default:
		fmt.Println("Unknown Shape")
	}
}

func main() {
	shapes := []Shapes {
			Circle{radius: 5},
			Rectangle{length: 3, width: 2},
			Square{length: 9},
			}
	
	// Print all shape dimension
	for i := 0; i < len(shapes); i++ {
		CalculateShapeCoverage(shapes[i])
	}
}

