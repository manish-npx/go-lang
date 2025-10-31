package main

import (
	"fmt"
	"slices" // Available from Go 1.21+
)

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

// ========== Integer Slice Example with Conversion ==========

// sliceArrExample demonstrates creation, appending, mutation, and manual conversion to []interface{}
func sliceArrExample() ([]int, []any) {

	// Create a slice with length 3, all values initialized to 0
	arr := make([]int, 3)

	// Assign values manually
	arr[0] = 20
	arr[1] = 30
	arr[2] = 50

	// Append elements dynamically
	arr = append(arr, 1)
	arr = append(arr, 2, 4)
	arr = append(arr, 1, 2, 3, 4, 5)

	// Mutate: multiply each element by 100
	for i, v := range arr {
		arr[i] = v * 100
	}

	// Convert []int to []interface{} manually (copy() won't work)
	newArr := make([]any, len(arr))
	for i, v := range arr {
		newArr[i] = v
	}

	return arr, newArr
}

// ========== String Slice Concepts ==========

func demoNilSlice() {
	var s []string // nil slice (no memory allocated)
	str := []string{}
	fmt.Println("Nill Slice", str, len(str))
	fmt.Println("Nil slice:", s, "Is nil?", s == nil, "Length == 0?", len(s) == 0)
}

func demoMakeAndAssign() []string {
	s := make([]string, 3) // allocates memory with len=3, cap=3
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("After assignment:", s, "len:", len(s), "cap:", cap(s))
	return s
}

func demoAppend(s []string) []string {
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("After append:", s)
	return s
}

func demoCopy(s []string) {
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy of slice:", c)
}

func demoSlicing(s []string) {
	sl1 := s[2:5] // slice from index 2 up to (not including) 5
	fmt.Println("Slice s[2:5]:", sl1)

	sl2 := s[:5] // from start to index 5 (excluded)
	fmt.Println("Slice s[:5]:", sl2)

	sl3 := s[2:] // from index 2 to end
	fmt.Println("Slice s[2:]:", sl3)
}

func demoLiteralDeclaration() {
	t := []string{"g", "h", "i"}
	fmt.Println("Literal slice:", t)
}

func demoSliceComparison() {
	t1 := []string{"g", "h", "i"}
	t2 := []string{"g", "h", "i"}

	if slices.Equal(t1, t2) {
		fmt.Println("t1 == t2: Slices are equal")
	} else {
		fmt.Println("t1 != t2: Slices are not equal")
	}
}

func demo2DSlice() {
	// Create a 2D slice (slice of slices)
	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2D slice:", twoD)
}

func demoSliceEle() {

	arrSlice := make([]int, 0)
	arrSlice = append(arrSlice, 10)
	arrSlice = append(arrSlice, 20)
	arrSlice = append(arrSlice, 30)
	arrSlice = append(arrSlice, 40)
	arrSlice = append(arrSlice, 50)
	arrSlice = append(arrSlice, 100, 200, 300, 400)
	fmt.Println(arrSlice)
}

// ========== MAIN ==========

func main() {

	fmt.Println("===== Array Integer Slice Example =====")
	demoSliceEle()
	fmt.Println("===== Integer Slice Example =====")
	arr, newArr := sliceArrExample()
	fmt.Println("Original int slice:", arr)
	fmt.Println("Converted []interface{} slice:", newArr)

	fmt.Println("\n===== Nil Slice Check =====")
	demoNilSlice()

	fmt.Println("\n===== Slice Creation and Assignment =====")
	s := demoMakeAndAssign()

	fmt.Println("\n===== Append Elements =====")
	s = demoAppend(s)

	fmt.Println("\n===== Copy Slice =====")
	demoCopy(s)

	fmt.Println("\n===== Slicing Slice =====")
	demoSlicing(s)

	fmt.Println("\n===== Slice Literal Declaration =====")
	demoLiteralDeclaration()

	fmt.Println("\n===== Slice Comparison =====")
	demoSliceComparison()

	fmt.Println("\n===== 2D Slice (Nested Slices) =====")
	demo2DSlice()
}
