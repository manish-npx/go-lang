package main

import (
	"fmt"
	"math"
)

// `const` declares a constant value.
const s string = "constant"

func main() {
	fmt.Println(s)

	// A `const` statement can appear anywhere a `var`
	// statement can.
	const n = 500000000 // arbitrary precision.
	const d = 3e20 / n
	const a = 10
	const b = a
	const dm = a / 20
	fmt.Println(a == b)
	fmt.Println(a == 10)
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	type Person struct {
		name    string
		age     int
		gender  string
		city    string
		state   string
		address string
		zip     int
	}

	var p = Person{
		name:    "Manish",
		age:     29,
		gender:  "male",
		city:    "Pune",
		state:   "Maharastra",
		address: "phase 1",
		zip:     400528,
	}

	fmt.Printf("Person name is %s", p.name)
	fmt.Println(" ")
	fmt.Println("Person name is ", p.name)

}
