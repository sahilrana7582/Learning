package main

import "fmt"

/*
	Concurrency:

	Concurrency is a way to handle multiple tasks at the same time.

	GoLang is a concurrent language.



*/

func printHello() {
	fmt.Println("Hello")
}

func main() {

	a := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			a <- i
		}()
	}

	for value := range a {
		fmt.Println(value)
	}

}
