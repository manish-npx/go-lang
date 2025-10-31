package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Task", id)

}

func main() {
	fmt.Println("coming from go Goroutines and wait group")

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i, &wg) //making goroutine by adding go keyword

		//anonymous function
		// func(i int) {
		// 	fmt.Println("i =>", i)
		// }(i)
	}

	wg.Wait()

}
