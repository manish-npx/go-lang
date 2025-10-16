package main

import {fmt,log}

func main() {
	name := "manish"
	fmt.Println(greet(name))
	fmt.Println("Hello goLang")
	log.Println("Hello GO")


}

func greet(person string) string {
	return "hi " + person
}
