package main

import (
	"fmt"
	"time"
)

func main() {
	dateTime := time.Now()

	// ✅ Format pattern uses reference time: 2006-01-02 15:04
	formatted := dateTime.Format("2006-01-02 15:04:05")

	fmt.Println("Current Date & Time:", formatted)

	year, month, day := dateTime.Date()
	hour, min, sec := dateTime.Clock()

	fmt.Printf("Year: %v, Month: %d, Day: %d\n", year, month, day)
	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n", hour, min, sec)
}
