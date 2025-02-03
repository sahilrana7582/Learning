package main

// This import is used when we want to import something from the other package in our project
// OR
// This import is used when we want to import something from the standard library
// GoLang has import the whole package from its standard library by default
// It not import not specific package, function, constant, expression like other programming language such as JavaScript, Python, Java etc.
import "fmt"

func main() {

	/*
		// There Are Four Types of Literals in GoLang
		// 1. Boolean
		// 2. Numeric
		// 3. String
		// 4. Rune
	*/

	/*
		// The Zero Value in GoLang
		// If We Define the variable without any value then it will be assigned with Zero Value
		// Zero Value of Boolean is False
		// Zero Value of Numeric is 0
		// Zero Value of String is ""
		// Zero Value of Rune is 0
		// var a int -> The Zero Value of int is 0
	*/

	var a int
	fmt.Println(a) // This will print 0 because the Zero Value of int is 0

	// The Zero Value of Boolean is False
	var b bool
	fmt.Println(b) // This will print False because the Zero Value of Boolean is False

	// The Zero Value of String is ""
	var c string
	fmt.Println(c) // This will print "" because the Zero Value of String is ""

	// The Zero Value of Rune is 0
	var d rune
	fmt.Println(d) // This will print 0 because the Zero Value of Rune is 0

	// In Golang there are so many ways to declare a variable
	// The First Way to declare a variable is using var keyword
	// The Second Way to declare a variable is using := keyword

	var a1, b1 int = 10, 10
	fmt.Println(a1, b1) // This will print 10 20 because the value of a is 10 and the value of b is 20

	// We can also define variable with different type

	var a2, b2, c2, d2 = 10, "Hello", 10.5, true
	fmt.Println(a2, b2, c2, d2) // This will print 10 Hello 10.5 true because the value of a is 10, the value of b is Hello, the value of c is 10.5 and the value of d is true

	/*
		//Shorthand Declaration with := Keyword
		// The Shorthand Declaration with := keyword is used to declare a variable with a value
		// GoLang will autoomatically detect the type of the variable
		// The Shorthand Declaration with := keyword is only used inside the function
		// The Shorthand Declaration with := keyword is not used outside the function
	*/
	a3 := 10

	fmt.Println(a3) // This will print 10 because the value of a is 10

	// We can also define variable with different type
	a4, b4, c4, d4 := 10, "Hello", 10.5, true
	fmt.Println(a4, b4, c4, d4) // This will print 10 Hello 10.5 true because the value of a is 10, the value of b is Hello, the value of c is 10.5 and the value of d is true

	// We can also assign the value to already  declared variable with := keyword, but the rule the type of value should be same as the type of variable
	a4 = 20 // The earlier value of a is 10 and now it is 20 both the type is int so it is  valid
	// a4 = "Hello" // This will give an error because the type of a is int and the type of value is string

	/*
		Const Keyword

		-> Const keyword is used to define a constant
		-> It is important to assign the value to the constant variable at the time of declaration
		-> We can't change the value of the constant variable
		-> We can't use the := keyword to declare the constant variable
		-> We can't use the var keyword to declare the constant variable

	*/

	const a5 = 20
	fmt.Println(a5) // This will print 20 because the value of a is 20
	//Const can hold only those values that compiler can evaluate at compile time.
	// We Cannot hold any value in the const variable that is evaluated or caculated at runtime.

	/*
		x := 10
		y := 20

		const z = x + y -> This will give error because the value of x and y is not known at compile time


		const a = 10
		cont b = 20

		const c = a + b -> This will work because the value of a and b is known at compile time because we have to give value to the constant while decaring the constant. So, it will be sure that value will be available at the compiler time.
	*/
}
