package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	email     string
}

func main() {

	/*
		what is the pointer?
		-> The Pointer is the variable that store the address of data of other value store in the memory.
		let assume we have a variable a that store the value 10
		a := 10
		// The Value of a is 10
		// The Address where the value of a is store in the memory is 0X1234322 -> This is called address
		So, the pointer is the variable that hold that address of the value

		to use pointer we have to use the & symbol
		-> &a -> This will return the address of the data of a.
	*/

	a := 1000

	pointerA := &a

	fmt.Println(pointerA)
	// Output will be something like this -> 0x14000104020
	/*
		The We Learned in the last function lession that if we pass the value to the function
		Then, it will make the copy first of that argument than pass it to the function
		So, if we make any change to the argument it will not affect the original value
	*/

	fun1(a)
	fmt.Println(a)
	// Output Will Be = 10, 1000

	// This time we are passing the memeory address of the value to the function
	// Sp, it will change the value to that address so it will change the value of the original value

	fun2(&a)
	fmt.Println(a)
	/*
		Output: 1000
		101
		101
	*/

}

// This function take the argument and try to change the value of the original value.
// But it will not change the value of the original value because it is pass the copy of the value to the function
func fun1(a int) {
	a = 10
	fmt.Println(a)
}

/*
	Whta if want to change the value of the original value
	then, we have to use the pointer of the variable only then we can change the value of the original value
*/

func fun2(a *int) {
	*a = 101
	fmt.Println(*a)
}
