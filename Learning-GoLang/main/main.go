package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	ch := make(chan string, 3) // Buffered channel with capacity of 3
	var wg sync.WaitGroup

	wg.Add(2) // Two producer goroutines

	// First producer goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			msg := "Hello | " + strconv.Itoa(i)
			fmt.Println("[Producer 1] Sending:", msg)
			ch <- msg                          // Sending message into channel
			time.Sleep(500 * time.Millisecond) // Simulate some work
		}
		fmt.Println("[Producer 1] Finished sending messages.")
	}()

	// Second producer goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			msg := "World | " + strconv.Itoa(i)
			fmt.Println("[Producer 2] Sending:", msg)
			ch <- msg                          // Sending message into channel
			time.Sleep(700 * time.Millisecond) // Different delay to mix up messages
		}
		fmt.Println("[Producer 2] Finished sending messages.")
	}()

	// Consumer goroutine
	go func() {
		for msg := range ch {
			fmt.Println("\t[Consumer] Received:", msg)
			time.Sleep(1 * time.Second) // Simulate processing time
		}
		fmt.Println("\t[Consumer] Channel closed. Stopping.")
	}()

	wg.Wait() // Wait for both producers to finish
	close(ch) // Close the channel after producers are done

	// Give some time for the consumer to finish processing
	time.Sleep(2 * time.Second)
	fmt.Println("Main function exiting.")
}
