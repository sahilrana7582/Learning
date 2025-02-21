package main

import (
	"fmt"
	"time"
)

/*
	Concurrency:

	Concurrency is a way to handle multiple tasks at the same time.

	GoLang is a concurrent language.



*/

func printHello(text string) {
	fmt.Println(text)
}

func main() {

	// go function function_name()
	// With the help go keyword in front of the function we can make the new go routine or we can say that function
	// will now run independenyly.

	// Now, the main function and new go routine function will run independently.
	// So, the main function will not wait for the new go routine function to finish.

	go printHello("Go Routine Function")
	fmt.Println("Main Function")

	/*
		Output:
		Main Function

		why, it is only print the main function only where is the go routine function?
		-> Because, the main function is finished excuting before the go routine function is finished.
		So, the main function will not wait for the go routine function to finish.
		So, the go routine function will not be able to print the output.
	*/

	// So, we have to make the main function to wait for the go routine function to finish.
	// We can do this by using the time.Sleep() function.
	// This function will make the main function sleep for given time and meanwhile the go routine function will run and done executing.
	// So, the main function will not finish before the go routine function is finished.

	time.Sleep(time.Millisecond * 100)

	/*
		Output:
		Main Function
		Go Routine Function
	*/

	// Next Topic ----> Channels
	// What are channels?
	// Channels are the way to communicate between the goroutines.
	// We can use the channels to send and receive the data between the goroutines.
	// We can use the channels to synchronize the goroutines.
	// Syntax: channel_name := make(chan type, size_of_channel) -> type can be any data type.
	// size_of_channel is the size of the channel.

	ch := make(chan string, 5)

	// Now, we can store 5 data in this channel.

}
