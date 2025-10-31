package main

import "fmt"

func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * fact(n-1)

}
func main() {
	var facto int = fact(3)
	fmt.Println("print facto", facto)

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println("FIB:", fib(5))

}
