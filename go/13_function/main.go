/* package main

import "fmt"

func sumOfThree(a, b, c int) int {
	return a + b + c
}

func sumOfTwo(a float64, b float64) float64 {
	return a + b
}
func varForSum(a, b, c int) {
	fmt.Println("Please enter a 1st number")
	fmt.Scan(&a)
	fmt.Println("Please enter a 2nd number")
	fmt.Scan(&b)
	fmt.Println("Please enter a 3rd number")
	fmt.Scan(&c)
}

func main() {
	var a, b, c int

	varForSum(a, b, c)
	if a+b+c == 0 {
		fmt.Println("Please enter valid number")
		varForSum(a, b, c)
	}
	sum := sumOfThree(a, b, c)
	fmt.Println("sum of 3 int", sum)

	floatSum := sumOfTwo(float64(a), float64(b))
	fmt.Println("sum of 2 float", floatSum)

}
*/

package main

import "fmt"

func sumOfThree(a, b, c int) int {
	return a + b + c
}

func sumOfTwo(a float64, b float64) float64 {
	return a + b
}

func directlyValueReturn() (int, int) {
	return 3, 7
}

func getValueFromInput(a, b, c *int) {
	fmt.Println("Please enter a 1st number")
	fmt.Scan(a)
	fmt.Println("Please enter a 2nd number")
	fmt.Scan(b)
	fmt.Println("Please enter a 3rd number")
	fmt.Scan(c)
}

func main() {
	var a, b, c int

	getValueFromInput(&a, &b, &c)
	x, y := directlyValueReturn()
	fmt.Printf("x value %v and y value is %v\n\n", x, y)

	sum := sumOfThree(a, b, c)
	fmt.Println("sum of 3 int", sum)

	floatSum := sumOfTwo(float64(a), float64(b))
	fmt.Println("sum of 2 float", floatSum)
}
