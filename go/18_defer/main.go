package main

import "fmt"

func sum(a, b int) int {
	return a + b
}
func demo() {
	//can not write something like that
	// func first(){
	// 	defer fmt.Println("first")
	// }
	// func second(){
	// 	defer fmt.Println("second")
	// }
	// func third(){

	// 	defer fmt.Println("third")
	// }

	first := func() {
		defer fmt.Println("first")
	}
	second := func() {
		defer fmt.Println("second")
	}
	third := func() {
		defer fmt.Println("third")
	}

	first()
	second()
	third()
}
func main() {
	add := sum(2, 4)
	demo()
	defer fmt.Println("defer example Adding:= ", add)
	defer fmt.Println("MSG!")
	fmt.Println("MSG@")
	defer fmt.Println("MSG#")

	//Logging
	f := withLogging()
	f()
}
func withLogging() func() {
	fmt.Println("Setup started")

	return func() {
		defer fmt.Println("Cleanup done")
		fmt.Println("Inside returned function")
	}
}
