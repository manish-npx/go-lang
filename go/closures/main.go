package main

import "fmt"

const counterString = "counter is"

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {

	count := counter()
	fmt.Println(counterString, count())
	fmt.Println(counterString, count())

	counts := counter()
	fmt.Println(counterString, counts())
	fmt.Println(counterString, counts())
}
