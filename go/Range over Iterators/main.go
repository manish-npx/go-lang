package main

import (
	"fmt"
	"iter"
)

// Create a simple iterator that counts 1 to 3
func CountToThree() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= 3; i++ {
			if !yield(i) { // send value to range loop
				return // stop if loop breaks
			}
		}
	}
}

func sumToNumbers() iter.Seq[int] {
	return func(yield func(int) bool) {
		sum := 0
		for i := 1; i < 5; i++ {
			sum += i
			if !yield(sum) {
				return
			}
		}
	}
}

func main() {
	// "Range over iterator"
	for num := range CountToThree() {
		fmt.Println(num)
	}

	for sumNum := range sumToNumbers() {
		fmt.Println("sum num is:= ", sumNum)
	}
}
