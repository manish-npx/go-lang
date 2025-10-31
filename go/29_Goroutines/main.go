package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Task", id)

}

func main() {
	fmt.Println("coming from go Goroutines")

	for i := 0; i < 10; i++ {

		go task(i) //making goroutine by adding go keyword

		//anonymous function
		func(i int) {
			fmt.Println("i =>", i)
		}(i)
	}

	time.Sleep(time.Second * 2)

}
