package main

import "fmt"

func main() {

	/*
		Composite Types -> Array
		-> Array is a collection of elements of the same type
		-> Array is a fixed size in GoLang

		-> Syntax -> var array_name [size]type
			* It is important to define the size of the array
			* It is important to define the type of the array

		-> The very important thing to remember while working with the arrray in the GoLang that
			the size of the array is also the type of the array

		-> To Understand this concept

			Let assumge we have two array
			var a [5]int
			var b [10]int

			Now, we have two array with a with size 5 and b with size 10

			Now, let's make new function that take array as input but
			func sum(a [5]int) int {}
			As, you can see that we can to define func is taking the aray of what length
			So, we can't pass the array of size 10 to this function
			So, we have to make different function for each size of the array
	*/

	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	//Other Way to define the array
	var b = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(b)
	// in this way we can define the array without defining the size of the array
	// it is flexible we can add as many element as we want during the time of declaration
	// but we can't add more element after the declaration

	fmt.Println(len(b))
	// len() is a function that return the length of the array

	/*
		-> In GoLang, mostly we use slice instead of array
		-> Slice is a dynamic array

		There are two ways to define the slice
			* var slice_name []type
			* make([]type, size)
		OR
		* var slice_name = []type{value1, value2, value3, ...}
		As You See we don't have to expicitly define the size of the slice

		-> We can also add new element to the slive even after the declaration
		-> We can also remove the element from the slice

		-> To add new element to the slice we can use append() function
			* There are three ways to append the slice with the new element

			* append(slice_name, value)
			//This way we can append the slice with the new but with one single element

			* append(slice_name, value1, value2, value3, ...)
			//This way we can append the slice with the new but with multiple element

			* append(slice_name, slice_name2...)
			//This way we can append the slice with the existing slice slice_name... is the variadic parameter which mean it return all the elements of the slice as one by one
				*For Eg:
					-> x := []int{1, 2, 3}
					x... => 1,2,3 this pattern return us all the valuee of the slice as one by one


		There is one important concept about the capacity of the slice-:
			-> len(slice_name) = this function return the length of the slice. The Number of the elements currently stored in the slice
			-> cap(slice_name) = this function return the capacity of the slice. The Number of the elements that the slice can store
			-> If the slice got full then it will first make the new slice of the double size of the current slice
			-> then, it will copy all the elements of the slice into the new slice

		The Other way to make the slice with the make() function
			-> There are two ways to use the make() function
				* make([]type_of_slice_you_want_to_make,    len_and_capacity_of_slice_you_want);
				* make([]type_of_slice_you_want_to_make,    len_of_slice,  capacity_of_slice);
	*/

	x1 := make([]int, 5)
	fmt.Println(x1)
	// It will return the new slice that has the 5 len and 5 capacity

	x2 := make([]int, 5, 10)
	fmt.Println(x2)
	// It will return the new slice with the length of 5 and the capacity of holding new element is 10

	var x3 = []int{1, 2, 3, 4, 5}
	fmt.Println(len(x3))
	// It will return the length of the current slice

	clear(x3)
	// This make the slice empty like this -> x3 = {}
	fmt.Println(len(x3))
	// Still the length of the slice will remain same as it was before

	x5 := []int{1, 2, 3, 4, 5}
	y5 := x5[1:3]
	fmt.Println(x5)
	fmt.Println(y5)
	y5[1] = 101

	fmt.Println(x5)

	fmt.Println(y5)

}
