package main

import "fmt"

// A simple generic function that swaps two values
func Swap[T any](a, b T) (T, T) {
	return b, a
}

//func printSLice(items []int) {
//func printSLice[T any](items []T) { //generic type by using any
//func printSLice[T string | int](items []T) { //generic type by using pipe | , we can use only 2 type other than string int
func printSLice[T comparable](items []T) { //generic type by using comparable, that support all the multiple type
	for _, item := range items {
		fmt.Println("Item ===> ", item)
	}
}

func main() {
	x, y := Swap[int](10, 20)
	fmt.Println(x, y) // 20, 10

	s1, s2 := Swap[string]("hello", "world")
	fmt.Println(s1, s2) // world, hello

	//==============================================

	nums := []int{1, 2, 3, 4, 5}
	lang := []string{"TS", "JS", "PHP", "GO"}
	boolVal := []bool{true, false, true}
	printSLice(nums)
	printSLice(lang)
	printSLice(boolVal)
}
