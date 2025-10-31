package main

import (
	"fmt"
	"math"
)

////////////////////////////////////////////////////////////////////////////////
// Step 1: Define the interface
////////////////////////////////////////////////////////////////////////////////

// geometry interface declares two method signatures.
// Any type with both methods (area() and perim()) implements this interface.
type geometry interface {
	area() float64
	perim() float64
}

////////////////////////////////////////////////////////////////////////////////
// Step 2: Define concrete types (structs)
////////////////////////////////////////////////////////////////////////////////

// rect represents a rectangle with width and height.
type rect struct {
	width, height float64
}

// circle represents a circle with a radius.
type circle struct {
	radius float64
}

////////////////////////////////////////////////////////////////////////////////
// Step 3: Implement interface methods for rect
////////////////////////////////////////////////////////////////////////////////

// area calculates and returns the area of the rectangle.
// Implements the area() method from geometry interface.
func (r rect) area() float64 {
	return r.width * r.height
}

// perim calculates and returns the perimeter of the rectangle.
// Implements the perim() method from geometry interface.
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

////////////////////////////////////////////////////////////////////////////////
// Step 4: Implement interface methods for circle
////////////////////////////////////////////////////////////////////////////////

// area calculates and returns the area of the circle.
// Implements the area() method from geometry interface.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// perim calculates and returns the perimeter (circumference) of the circle.
// Implements the perim() method from geometry interface.
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

////////////////////////////////////////////////////////////////////////////////
// Step 5: Use the interface as a function parameter
////////////////////////////////////////////////////////////////////////////////

// measure accepts any type that satisfies the geometry interface.
// It prints the object, its area, and perimeter.
func measure(g geometry) {
	fmt.Println("Measuring shape:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
}

////////////////////////////////////////////////////////////////////////////////
// Step 6: Use type assertion to detect specific type
////////////////////////////////////////////////////////////////////////////////

// detectCircle checks whether the given geometry is a circle using type assertion.
// If it is a circle, it prints its radius.
func detectCircle(g geometry) {
	// g.(circle) is a type assertion; 'ok' is true if g is actually a circle
	if c, ok := g.(circle); ok {
		fmt.Println("Detected a circle with radius:", c.radius)
	} else {
		fmt.Println("Not a circle.")
	}
}

////////////////////////////////////////////////////////////////////////////////
// Step 7: Main function — demonstrate usage
////////////////////////////////////////////////////////////////////////////////

func main() {
	// Create a rectangle and a circle
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Call measure() with both shapes — polymorphism in action
	fmt.Println("---- Measuring Rectangle ----")
	measure(r)

	fmt.Println("\n---- Measuring Circle ----")
	measure(c)

	// Try to detect if each shape is a circle using type assertion
	fmt.Println("\n---- Type Assertion Checks ----")
	detectCircle(r) // Will not match
	detectCircle(c) // Will match and print radius
}
