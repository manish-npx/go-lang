package main

import (
	"fmt"
)

type Person struct {
	name   string
	email  string
	age    int
	mobile int
}

func main() {
	var person Person

	fmt.Println("Enter name")
	fmt.Scan(&person.name)
	fmt.Println("Enter email")
	fmt.Scan(&person.email)
	fmt.Println("Enter age")
	fmt.Scan(&person.age)
	fmt.Println("Enter mobile")
	fmt.Scan(&person.mobile)

	res := map[string]any{
		"name":   person.name,
		"email":  person.email,
		"age":    person.age,
		"mobile": person.mobile,
	}

	for i, v := range res {
		fmt.Printf("Person %v is %v\n", i, v)
	}

	resSlice := make([]any, 0)
	resSlice = append(resSlice, person.name, person.mobile, person.email, person.age)

	fmt.Println("Person Value is ", res)
	fmt.Println("Person Value is ", resSlice)
}
