package main

import (
	"fmt"
	"sync"
)

func processNum(numChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Process Num", <-numChan)

}

func processMsg(msgChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-msgChan
	fmt.Println("Processing ", msg)

}

func taskOnlyChanel(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Processing via only channel")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	numChan := make(chan int)

	go processNum(numChan, &wg)

	numChan <- 5

	//Channel syntax
	messageChan := make(chan string)
	go processMsg(messageChan, &wg)
	messageChan <- "Ping Pong"

	wg.Wait()

	/* without wait */
	done := make(chan bool)

	go taskOnlyChanel(done)
	<-done
	//without wait group channel
}
