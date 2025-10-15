package main

import "fmt"

func main() {
	name := "manish"
	fmt.Println(greet(name))
	fmt.Println("Hello goLang")

}

func greet(person string) string {
	return "hi " + person
}
