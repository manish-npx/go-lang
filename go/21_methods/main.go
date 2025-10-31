package main

import "fmt"

type rect struct {
	height, width int
}

// area uses a value receiver because it just reads the struct fields
func (r rect) area() int {
	return r.height * r.width
}

// scale uses a pointer receiver because it modifies the struct
func (r *rect) scale(factor int) {
	r.height = r.height * factor
	r.width = r.width * factor
}

func main() {
	// Create a rect struct value
	r := rect{height: 10, width: 20}

	// Call area() — value receiver method
	fmt.Println("Original area:", r.area()) // 200

	// Call scale() — pointer receiver method (modifies the struct)
	r.scale(2)

	// Call area() again after scaling
	fmt.Println("Scaled area:", r.area()) // 800
}
