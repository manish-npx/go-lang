package main

import "fmt"

func MapMethodExample() map[string]int {
	m := make(map[string]int)
	m["age1"] = 20
	m["age2"] = 30
	age := m["age1"]
	fmt.Printf("age value is  %v\n", age)
	return m
}

//If key value does not exist in the map then return the zero value in the case int zero string empty other this return nil
func MapMethodsInterface() {
	m := make(map[string]any)
	m["name"] = "manish"
	m["age"] = 100
	m["address"] = "Pune"
	Person := map[string]any{
		"name":    "manish",
		"age":     100,
		"address": "Pune",
	}
	fmt.Println("Person", Person)
	deleteMapValue(Person, "name")

}

func deleteMapValue(m map[string]any, v string) {
	delete(m, v)
	fmt.Printf("Person %v deleted Successfully\n", v)
}

func main() {
	res := MapMethodExample()
	MapMethodsInterface()
	fmt.Println("Map Methods are", res)
}
