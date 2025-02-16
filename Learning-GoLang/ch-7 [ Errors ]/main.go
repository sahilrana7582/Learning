package main

import (
	"errors"
	"fmt"
)

/*
	Error Chapeter:
	In This Chapter We Will Learn About Errors In Go, how to handle them, how to create our own errors, and how to use the errors package.s


	When function excute successfully, it returns a nil value.

	nil mean No Error.

	While, function did something wrong, it return the value of type error.

	Then, we have to catch that error value, handle that value.


	error, value := function() // Assume, our functiom return two values. Error and Value.

	if error != nil {
		// Handle the Error Logic;
	}

	fmt.Println(value) // Everything is fine, we got the value.
*/

// Create a custom error

func divideEven(a int) (string, error) {
	if a%2 == 0 {
		return "Even Number", nil
	}

	return "", errors.New("Odd Number")
}

func main() {

	val, err := divideEven(2)
	if err != nil {
		fmt.Println("Error Occured-> ", err)
	}

	fmt.Println(val)

	val, err = divideEven(3)

	if err != nil {
		fmt.Println("Error Occured-> ", err)
	}

	fmt.Println(val)
}
