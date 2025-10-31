package main

import "fmt"

/*
	__________________________________________________________________________________________________
	| Syntax                 | Type    | Size     | Resizable? | Passed by value?                     |
	|------------------------|---------|----------|------------|--------------------------------------|
	| []string{"a"}          | Slice   | Dynamic  | Yes        | No (header copied, data shared)      |
	| [...]string{"a"}       | Array   | Fixed    | No         | Yes (entire array is copied)         |
	| arr := [3]int{}        | Array   | Fixed    | No         | Yes                                  |
	| slice := arr[1:3]      | Slice   | Dynamic  | Yes        | No (references array underneath)     |
	| make([]int, 5)         | Slice   | Dynamic  | Yes        | No                                   |
	|_________________________________________________________________________________________________|

*/

func FixedSizeArray() [5]int {
	numbers := [5]int{10, 20, 30, 40, 50}
	return numbers
}

func InferredSizeArray() [3]string {
	labels := [...]string{"red", "green", "blue"}
	return labels
}

func SliceLiteralExample() []string {
	animals := []string{"cat", "dog", "bird"}
	return animals
}

func SliceFromArrayExample() []int {
	arr := [5]int{10, 20, 30, 40, 50}
	subSlice := arr[1:4]
	subSlice[1] = 300
	fmt.Println("Original array after slice modification:", arr)
	return subSlice
}

func MakeSliceExample() []int {
	numbers := make([]int, 5)
	numbers[0] = 100
	return numbers
}

func DynamicSliceGrowth() []string {
	words := []string{}
	words = append(words, "hello")
	words = append(words, "world")
	return words
}

func main() {
	fmt.Println("1. Fixed Size Array        	:", FixedSizeArray())
	fmt.Println("2. Inferred Size Array     	:", InferredSizeArray())
	fmt.Println("3. Slice Literal           	:", SliceLiteralExample())
	fmt.Println("4. Slice From Array        	:", SliceFromArrayExample())
	fmt.Println("5. Slice Created with make		:", MakeSliceExample())
	fmt.Println("6. Growing Slice with append	:", DynamicSliceGrowth())
}
