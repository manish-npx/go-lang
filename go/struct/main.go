package main

import "fmt"

type Person struct {
	name  string
	age   int
	city  string
	state string
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 25          //adding default value
	p.city = "varanasi" //adding default value
	p.state = "UP"      //adding default value
	return &p
}

func main() {
	person := Person{
		name:  "Manish",
		age:   20,
		city:  "pune",
		state: "Mah",
	}
	p := newPerson("Deep")
	fmt.Println(newPerson("Jon")) // Create person using constructor function, returns pointer

	// Define and initialize an anonymous struct (no explicit type name)
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println("Person Name", person.name)
	fmt.Println("Person Age", person.age)
	fmt.Println("Person city", person.city)
	fmt.Println("Person state", person.state)
	fmt.Println("====================================")
	fmt.Println("DOG", dog)
	fmt.Println(&Person{name: "Ann", age: 40}) // Pointer to struct literal
	fmt.Println("====================================")
	fmt.Println("Person Name", p.name)
	fmt.Println("Person Age", p.age)
	fmt.Println("Person city", p.city)
	fmt.Println("Person state", p.state)

}
