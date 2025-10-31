package main

import "fmt"

// Worker function: reads from channel until it's closed
func emailSending(emailSending <-chan string, done chan<- bool) {
	for email := range emailSending { // automatically ends when channel closed
		fmt.Println("📧 Email is:", email)
	}
	done <- true // worker signals it finished
}

func main() {
	emailSendingChan := make(chan string, 20)
	doneChan := make(chan bool)

	numWorker := 3

	// Start 3 worker goroutines
	for i := 0; i < numWorker; i++ {
		go emailSending(emailSendingChan, doneChan)
	}

	// Producer: send emails into channel
	emails := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, mail := range emails {
		emailSendingChan <- fmt.Sprintf("%d_@mail.com", mail)
	}

	// ❗ Close the channel to signal no more data
	close(emailSendingChan)

	// ✅ Wait for all workers to finish
	for i := 0; i < numWorker; i++ {
		<-doneChan
	}

	fmt.Println("🎉 All emails processed successfully!")
}
