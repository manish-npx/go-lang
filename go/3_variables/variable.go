package main

import "fmt"

func Variable(varKey string) {
	fmt.Println(varKey)

}

var variableKey, manish string = "new Variable", "manish"
var num1, num2, num3 int = 1, 2, 3 //multiple declare in one line
var b1 int = 50                    //function scope (func)

func VarCheck(a string, b int, c int, d bool, f string) {
	const TRUE = true
	name := "manish" // block scope
	a1 := 20         //short hand declare means var a int = 20
	if a1 == b1 && TRUE {
		fmt.Println("func variable and var variable value are same ")
	} else {
		fmt.Println("not same")
	}
	Variable(variableKey)
	fmt.Println("Hello World"+" "+name, manish)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
}
