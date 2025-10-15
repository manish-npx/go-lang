package main

import "fmt"

func main() {

	const TRUE bool = true
	const FALSE bool = false

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(TRUE && FALSE)
	fmt.Println(TRUE || FALSE)
	fmt.Println(!TRUE)
}
