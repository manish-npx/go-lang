package main

import (
	"fmt"
	"time"
)

func NormalSwitch() {
	week := "mon"

	switch week {
	case "mon":
		fmt.Println("Monday")
	case "tue":
		fmt.Println("Tuesday")
	}
}

func TraditionalSwitch() {
	i := 2
	switch i {
	case 1:
		fmt.Println("No is", i)
	case 2:
		fmt.Println("No is", i)
	default:
		fmt.Println("Not found No")
	}
}

func DateSwitch() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Enjoy weekend")
	default:
		fmt.Println("Working days")
	}
}

func TimeSwitch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon", t.Hour())
	default:
		fmt.Println("It's after noon", t.Hour())
	}
}

func TypeSwitch() {
	checkVarType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Boolean type :", t, i)
		case int:
			fmt.Println("Integer Type: ", t, i)
		case string:
			fmt.Println("String Type: ", t, i)
		default:
			fmt.Println("No Types :", t)
		}
	}
	checkVarType(true)
	checkVarType(20)
	checkVarType("manish")
}

func main() {
	i := 20
	// Can't do this outside type switch, remove this line:
	// t := i.(type)

	if i != 20 {
		NormalSwitch()
	} else {
		TraditionalSwitch()
	}

	DateSwitch()
	TimeSwitch()
	TypeSwitch()

	// Print type of i in another way:
	fmt.Printf("Type of i is %T\n", i)
}
