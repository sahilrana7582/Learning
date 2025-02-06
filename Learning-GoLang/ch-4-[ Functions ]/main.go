package main

import "fmt"

func main() {
	/*
		In This Chapter We will learn about the functions in the GoLang
	*/

	// These are function that we defined below
	fun1()
	fun2(121)
	fun3(101, "King-Sahil", true, 1212)

	// The Limitation of the GoLang function is that we have to pass the all argument that be used during the function definition
	// In other programming language like JavaScript we can have the power of optional argument
	/*
		function functin_name(a, b, c, d?){
			D is optional in the JavaScript
		}
	*/

	/*
		Look at the Line Number #80
	*/

	a := Person{
		firstName: "Sahil",
		lastName:  "Rana",
	}

	b := Person{
		firstName: "Sahil",
		age:       23,
	}

	fun4(a) // This is the only way to work with the optional argument in the GoLang
	fun4(b) // This is the only way to work with the optional argument in the GoLang

	fmt.Println(fun5(3, 6, 7, 8))
	fmt.Println(fun5(3, 6, 7, 8, 9, 10, 10, 10))
	fmt.Println(fun5(3, 6, 7, 8, 1, 1, 1, 1, 1))
	fmt.Println(fun5(3, 6, 7, 8, 1, 1, 1))
	fmt.Println(fun5(3))

	/*
		Anonymous Function
		-> These are function that don't have anu name
		-> These are function that are used to define the function inside the function
	*/

	f := func(a, b int) {
		fmt.Println("Printing the value from the anonymous function", a, b)
	}
	f(10, 20)

	// We Can also convert them into IIFE Immediately Invoked Function Expression
	// IIFE
	for i := 0; i < 5; i++ {
		func(a int) {
			fmt.Println("Printing the value from the anonymous function", a)
		}(i)
	}

	/*
		Closure Function
		-> These are function that are used to define the function inside the function
		-> These function can access the variable of the both inner as well as outer function
	*/

	x := 1000

	f1 := func() {
		fmt.Println("Printing From the closure function", x)
		// Here we have the access of the outer variable x, It mean we can also change that too
		x = 9000
	}
	f1()
	fmt.Println(x)
	// Output will be : 1000 -> 9000

	/*
		Let Learn about the Defer Keyword
		-> This keyword is used with the function
		-> This keyword is used to defer the execution of the function
		Syntax:
		defer function_name()
		-> Now this function will only run when the main function will just about the return the control



		Look at the line number #180 for example
	*/

	fun6()
}

/*
	Following we will learn how to declare the function in the GoLang
	-> Syntac of the function is below:
	we declare the function with the help of the keyword ---->func<----
		func function_name(){
			//Function Body

			// This Function will take no argument and return no value
		}

		func function_name(argument_name argument_type){
			//Function Body

			//This function will take one argument and no return value
		}

		func function_name(argument_name argument_type1, argument_name2 arguement_type2) return_type {
			//Function Body


			//This function will take two argument and return a value
		}


		Important Points to remember while working with the function int tge GoLang:
		-> We can pass as many argument as we want to pass in the function
		-> We can also return as many values as we want to return from the function
		Eg:
			func function_name(arg1, type, arg2 type, arg3 type,..........) int, int, int, string, string, bool,........{
				//Here we can return as many values from the functions as we want.
				//We can return as many values as we want
			}

*/

func fun1() {
	fmt.Println("This Function is only used for the printing the message")
}

func fun2(a int) {
	fmt.Println("This function is taking the one argument and no return value. We can use that value in this function body", a)
}

func fun3(a int, b string, c bool, d int) (int, string, bool, error) {
	fmt.Printf("This function is taking the four argument and return four value. We can use that value in this function body. The value of a is %d, the value of b is %s, the value of c is %t, the value of d is %d\n", a, b, c, d)
	return a, b, c, nil
}

/*
	If Want to use this optional argument we have to do it with the help of the struct
*/

type Person struct {
	firstName string
	lastName  string
	age       int
}

func fun4(p Person) {
	fmt.Println("This function is taking the one argument and no return value. We can use that value in this function body", p)
}

func fun5(a int, b ...int) []int {

	/*
		...int
		This pattern is used to store all the extra argument in the function as the form of the slice
	*/
	ans := make([]int, 0, len(b))
	for _, v := range b {
		ans = append(ans, v)
	}
	return ans
}

// Defer

func fun6() {
	/*
		As You Can see we are calling the fun7() function in the beginning of the fun6()
		But it will still print its result only after the when the func6() will fininsh its all statment inside it
	*/
	defer fun7()
	for i := 1; i < 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println("Func 6 is just about to finish. The Defered Function will print its result now")
}

func fun7() {
	fmt.Println("I am the Deferred function 7")
}
