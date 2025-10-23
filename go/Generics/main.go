package main

import "fmt"

// A simple generic function that swaps two values
func Swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	x, y := Swap[int](10, 20)
	fmt.Println(x, y) // 20, 10

	s1, s2 := Swap[string]("hello", "world")
	fmt.Println(s1, s2) // world, hello
}
