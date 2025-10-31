package main

import (
	"fmt"
	"sync"
)

// processNum receives a slice of integers from a channel and processes each element
func processNum(numChan <-chan []int, wg *sync.WaitGroup) {
	// Mark this goroutine as done once processing is complete
	defer wg.Done()

	// Receive one slice from the channel
	for _, val := range <-numChan {
		fmt.Println("Processing number:", val)
	}
}

// processMsg receives and prints a message from a channel
func processMsg(msgChan <-chan string, wg *sync.WaitGroup) {
	// Mark this goroutine as done after execution
	defer wg.Done()

	// Receive and print message
	fmt.Println("Processing message:", <-msgChan)
}

func main() {
	// WaitGroup is used to wait for multiple goroutines to finish
	var wg sync.WaitGroup

	// We will wait for two goroutines (processNum + processMsg)
	wg.Add(2)

	// Create channels
	// Unbuffered channels block until a receiver is ready
	numChan := make(chan []int)
	msgChan := make(chan string)

	// Sample data (slice of integers)
	sliceChan := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Launch goroutines for processing
	go processNum(numChan, &wg)
	go processMsg(msgChan, &wg)

	// Send data into channels
	// Since the channels are unbuffered, these sends will block
	// until the respective goroutines are ready to receive
	msgChan <- "Ping Pong"
	numChan <- sliceChan

	// Close channels to indicate no more data will be sent
	close(numChan)
	close(msgChan)

	// Wait for all goroutines to finish
	wg.Wait()

	/* buffer example  */

	fmt.Println("✅ All goroutines completed successfully!")
	fmt.Println("✅ Buffer example !")

	ch := make(chan string, 1)
	ch <- "Hello"
	fmt.Println(<-ch)
}
