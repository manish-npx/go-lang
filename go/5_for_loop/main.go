package main

import "fmt"

// ForLoop demonstrates the classic C-style for loop.
func ForLoop() {
	for i := 0; i <= 4; i++ {
		fmt.Println("For Loop - i value is:", i)
	}
}

// WhileLoop shows how to implement a while-style loop using 'for'.
func WhileLoop() {
	i := 1
	for i <= 3 {
		fmt.Println("While Loop - i value is:", i)
		i++
	}
}

// InfiniteLoop demonstrates an infinite loop with a break.
func InfiniteLoop() {
	for {
		fmt.Println("Infinite Loop (with break)")
		break
	}
}

// RangeLoop demonstrates looping over a slice using 'range'.
func RangeLoop() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("Range Loop over Slice:")
	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}
}

// LoopOverMap demonstrates looping over a map using 'range'.
// map[string]interface{} is used to allow mixed value types (e.g., string and int).
func LoopOverMap() {
	numbers := map[string]interface{}{
		"name":  "Manish",
		"age":   20,
		"city":  "Pune",
		"state": "Amhara",
	}

	fmt.Println("Loop Over Map:")
	for key, value := range numbers {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
}

// LoopOverString demonstrates iterating over a string (rune by rune).
func LoopOverString() {
	const str = "Hello World"

	fmt.Println("Loop Over String (runes):")
	for i, ch := range str {
		fmt.Printf("Index: %d, Rune: %c\n", i, ch)
	}
}

// LoopOverChannel demonstrates how to send and receive values using a channel.
func LoopOverChannel() {
	// 'make' is a built-in function used to allocate and initialize:
	// - Slices
	// - Maps
	// - Channels (as in this example)
	ch := make(chan int)

	// Launch a goroutine to send data to the channel
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i // send value into channel
		}
		close(ch) // Close channel to signal that no more values will be sent
	}()

	// Receive data from the channel until it is closed
	fmt.Println("Loop Over Channel:")
	for val := range ch {
		fmt.Println("Received value:", val)
	}
}

// main is the entry point of the program.
func main() {
	fmt.Println("=== For Loop ===")
	ForLoop()

	fmt.Println("=== While Loop ===")
	WhileLoop()

	fmt.Println("=== Infinite Loop ===")
	InfiniteLoop()

	fmt.Println("=== Range Loop over Slice ===")
	RangeLoop()

	fmt.Println("=== Loop Over Map ===")
	LoopOverMap()

	fmt.Println("=== Loop Over String ===")
	LoopOverString()

	fmt.Println("=== Loop Over Channel ===")
	LoopOverChannel()
}
