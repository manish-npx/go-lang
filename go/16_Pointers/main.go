package main

import "fmt"

// lbl1: This function takes an int by value; changes inside do NOT affect caller
func zeroval(ival int) {
	ival = 0 // changes local copy only
}

// lbl2: This function takes a pointer to int; changes inside affect caller
func zeroptr(iptr *int) {
	*iptr = 0 // dereference pointer and change actual value
}

// lbl3: This function returns the pointer to an int variable
func getPointerToInt(val int) *int {
	return &val // returns pointer to local copy (dangerous outside this function)
}

// lbl4: This function swaps two ints by using pointers
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func pointerExample(a, b int) {

	x := a
	var y *int
	y = &x
	z := &b
	fmt.Println("x printing", x, y, z)
}

func main() {
	i := 1
	j := 2

	fmt.Println("initial i:", i)
	fmt.Println("initial j:", j)

	pointerExample(i, j)

	zeroval(i)
	fmt.Println("after zeroval(i):", i) // unchanged, passed by value

	zeroptr(&i)
	fmt.Println("after zeroptr(&i):", i) // changed to 0, passed by pointer

	fmt.Println("address of i:", &i) // prints memory address of i

	// swap i and j using pointers
	swap(&i, &j)
	fmt.Println("after swap(&i, &j): i =", i, "j =", j)
}
