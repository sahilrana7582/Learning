// package main

// import (
// 	"fmt"
// 	"time"
// )

// /*
// 	Concurrency:

// 	Concurrency is a way to handle multiple tasks at the same time.

// 	GoLang is a concurrent language.

// */

// // func printHello(text string) {
// // 	fmt.Println(text)
// // }

// func main() {

// 	// fmt.Println("--------------------- Go Routines ---------------------")
// 	// fmt.Println("")
// 	// fmt.Println("")

// 	// // go function function_name()
// 	// // With the help go keyword in front of the function we can make the new go routine or we can say that function
// 	// // will now run independenyly.

// 	// // Now, the main function and new go routine function will run independently.
// 	// // So, the main function will not wait for the new go routine function to finish.

// 	// go printHello("Go Routine Function")
// 	// fmt.Println("Main Function")

// 	// /*
// 	// 	Output:
// 	// 	Main Function

// 	// 	why, it is only print the main function only where is the go routine function?
// 	// 	-> Because, the main function is finished excuting before the go routine function is finished.
// 	// 	So, the main function will not wait for the go routine function to finish.
// 	// 	So, the go routine function will not be able to print the output.
// 	// */

// 	// // So, we have to make the main function to wait for the go routine function to finish.
// 	// // We can do this by using the time.Sleep() function.
// 	// // This function will make the main function sleep for given time and meanwhile the go routine function will run and done executing.
// 	// // So, the main function will not finish before the go routine function is finished.

// 	// time.Sleep(time.Millisecond * 100)

// 	// /*
// 	// 	Output:
// 	// 	Main Function
// 	// 	Go Routine Function
// 	// */

// 	// // Next Topic ----> Channels
// 	// // What are channels?
// 	// // Channels are the way to communicate between the goroutines.
// 	// // We can use the channels to send and receive the data between the goroutines.
// 	// // We can use the channels to synchronize the goroutines.
// 	// // Syntax: channel_name := make(chan type, size_of_channel) -> type can be any data type.
// 	// // size_of_channel is the size of the channel.

// 	// ch := make(chan string, 5)

// 	// fmt.Println("")
// 	// fmt.Println("")
// 	// fmt.Println("")
// 	// fmt.Println("")

// 	// fmt.Println("--------------------- Channel ---------------------")
// 	// fmt.Println("")
// 	// fmt.Println("")

// 	// // Now, we can store 5 data in this channel.
// 	// // how to send the data to the channel?
// 	// // channel_name <- data -> This way we can send the data to the channel.
// 	// // How to receive the data from the channel?
// 	// // data := <- channel_name -> This way we can receive the data from the channel.

// 	// ch <- "Hello From the Channel"

// 	// fmt.Println(<-ch)
// 	// Output: Hello From the Channel

// 	// Things to remember:
// 	// 1. A single go routine rarely use both operations of sending and receiving data with the same channel.
// 	// 2. So, we have to make our channel readFromChannel and writeToChannel as separate go routines.

// 	// While passing channels as parameters, we have to pass them as pointers.
// 	// function_nam(channel_name <- chan type) -> This was we can pass the channel as a read only channel.
// 	// function_nam(channel_name chan <- type) -> This was we can pass the channel as a write only channel.

// 	// ch1 := make(chan string, 5)

// 	// go writeOnlyChannels(ch1)

// 	// time.Sleep(time.Second * 1)
// 	// // go readOnlyChannels(ch1)

// 	// for val := range ch1 {
// 	// 	fmt.Println(val)

// 	// 	// Comment the line in writeOnlyChannels close(ch1) function to see the error
// 	// 	// it will give an error because with range for loop will keep iterating on the channel looking for the next value
// 	// 	// So, after iterating on the channel it will try read the value from the channel but there is no value in the channel
// 	// 	// so, it will give an error.
// 	// 	// So, we have to close the channel after we are done with the channel.
// 	// 	// close(ch1)
// 	// 	// this way with range loop the channel will know that the channel is closed and it will not try to read further values from the channel.
// 	// }

// 	// for len(ch) > 0 {
// 	// 	fmt.Println(<-ch + " | <- Read Only Channel")
// 	// it is not the case with for loop because for loop will iterate on the channel and it will not try to read the value from the channel.
// 	// }

// 	chan1 := make(chan int)
// 	chan2 := make(chan int)

// 	go func() {
// 		time.Sleep(time.Second * 7)
// 		for i := 0; i < 5; i++ {
// 			chan1 <- i
// 		}
// 		close(chan1)
// 	}()

// 	go func() {
// 		time.Sleep(time.Second * 7)
// 		for i := 0; i < 5; i++ {
// 			chan2 <- i
// 		}
// 		close(chan2)
// 	}()

// 	select {
// 	case val, ok := <-chan1:
// 		if ok {
// 			fmt.Printf("Hello From Channel 1: %d\n", val)
// 		}
// 	case val, ok := <-chan2:
// 		if ok {
// 			fmt.Printf("Hello From Channel 2: %d\n", val)
// 		}
// 	case <-time.After(time.Second * 5):
// 		fmt.Println("Timeout")
// 	}

// }

// // func readOnlyChannels(ch <-chan string) {
// // 	for len(ch) > 0 {
// // 		fmt.Println(<-ch + " | <- Read Only Channel")
// // 		// ch <- "Hello From Read Only Channel"
// // 		// uncomment this line to see the error
// // 		// because we can't write to a read only channel
// // 	}
// // }

// // func writeOnlyChannels(ch chan<- string) {
// // 	for i := 0; i < 5; i++ {
// // 		ch <- "Hello From | -> " + strconv.Itoa(i)
// // 		fmt.Println("Hello From Write Only Channel")

// // 		// fmt.Print(<-ch)
// // 		// uncomment this line to see the error
// // 		// because we can't read from a write only channel
// // 	}
// // 	close(ch)
// // }

// // Go Routines Cancelation
// // Go Routines Cancelation

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func fetchData(ctx context.Context, url string) {
// 	// Simulating a network request
// 	select {
// 	case <-time.After(3 * time.Second): // Simulate delay
// 		fmt.Println("Fetched data from:", url)
// 	case <-ctx.Done(): // Context canceled (either timeout or explicit cancellation)
// 		fmt.Println("Request to", url, "canceled:", ctx.Err())
// 	}
// }

// func main() {
// 	// Creating a cancelable context with a 2-second timeout
// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 	defer cancel() // Ensure cancel is called to release resources

// 	// Start a goroutine to fetch data
// 	go fetchData(ctx, "https://example.com")

// 	// Wait for the context to be done (either through timeout or cancellation)
// 	<-ctx.Done()                // Block until the context is done (i.e., goroutine gets canceled)
// 	time.Sleep(1 * time.Second) // Give some time for the goroutine to finish
// 	fmt.Println("Main goroutine exiting")
// }

package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context, id int) {
	select {
	case <-time.After(time.Duration(15) * time.Second): // Simulate work
		fmt.Printf("Task %d completed\n", id)
	case <-ctx.Done(): // Listen for cancellation
		fmt.Printf("Task %d canceled: %s\n", id, ctx.Err())
	}
}

func main() {
	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Launch multiple tasks
	for i := 1; i <= 3; i++ {
		go task(ctx, i)
	}

	// Wait for the context to be canceled (timeout or manual cancel)
	time.Sleep(5 * time.Second) // Simulate some work in the main goroutine
	cancel()

	if val, ok := <-ctx.Done(); ok {
		fmt.Println("Context canceled:", val)
	} else {
		fmt.Println("Context not canceled")
	}

	fmt.Println("Main goroutine exiting")
}
