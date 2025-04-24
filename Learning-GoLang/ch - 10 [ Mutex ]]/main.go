package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int = 0
	var lock sync.Mutex

	increament := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Println("Increamented count to: ", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Println("Decremented count to: ", count)
	}

	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increament()
		}()
	}

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement()
		}()
	}

	wg.Wait()
	fmt.Println("Final count: ", count)
	// The final count will be 0 because we are incrementing and decrementing the count in the same go routine
}
