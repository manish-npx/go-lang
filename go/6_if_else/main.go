package main

import "fmt"

func main() {
	var Age int = 10
	checkNum := 20
	if Age > checkNum {
		fmt.Println("You can vote and your age is ", Age)
	} else {
		fmt.Println("You can't")
	}

	sum := func(a, b int) int {
		return a + b
	}

	if sum(1, 0) != 0 {
		fmt.Println("Sum is ", sum(1, 0))
	} else {
		fmt.Println("Sum is 0")
	}
}
