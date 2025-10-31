package main

import "fmt"

func stringExample() {
	s := "Hello, 世界" // A string containing both ASCII and non-ASCII (Unicode) characters

	fmt.Println("Original string:", s)

	fmt.Println("\nByte-wise iteration (UTF-8 encoding):")

	for i, v := range s {
		var r rune = rune(v)
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n, Char:%d", i, v, s[i], r)
		fmt.Printf("Byte Index: %d, Rune: %c, Unicode: %U, Decimal: %d\n", i, r, r, r)
	}

}

func runesExample(r rune) {

	fmt.Println("Runes is =", r)
}

func main() {
	const a = 'A' //runes
	stringExample()
	runesExample(a)
}
