package main

import (
	"fmt"
	"maps"
)

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

	/*
		To slice the slice we can use the following syntac:
		* slice_name[start_index:end_index]
		* start_index = the index of the element from where we want to start the slice
		* end_index = the index of the element till where we want to slice the slice
		* The end_index is not included in the slice
		* The start_index is included in the slice
		For Eg:
			x := []int{1, 2, 3, 4, 5}
			y := x[0:3] // This will Return the slice {1, 2}
			//3 will not included in the slice
	*/
	x5 := []int{1, 2, 3, 4, 5}
	y5 := x5[1:3]
	fmt.Println(x5)
	fmt.Println(y5)

	/*
		The Important thing to remember is that:
		 * If we get the slice from the another slice and then we try to change into the new slice then it will change the original slice
		 * Same we we will change the new slice then it will also change the original slice as well
		 As you can see in the following example
	*/
	y5[1] = 101 // This will change the original slice as well
	fmt.Println(x5)
	fmt.Println(y5)

	/*
		-----------------------------------------------
			MAPS
			Now, let talk about the maps:
			Map is the data structure that will store the data with respect to the key
			Map is the collection of key-value pair

			There are three ways to define the maps:
			* var map_name [Key_Name]Type_of_Value_You_Want_To_Store
				-> Type can be any type like int, string, float, bool, etc.
			* map_name := map[Key_Name_Type]Type_of_Value_You_Want_To_Store{}
			* map_name := map[Key_Name_Type]Type_Of_Value_You_Want_To_Store{
				Key_Name_1 : value_1,
				Key_Name_2 : value_2,
				Key_Name_3 : value_3,
			}
	*/

	mpp := map[string]int{
		"Sahil": 999,
		"King":  79,
	}

	fmt.Println(mpp)

	//How to access the data from the the map
	// map_name[Key_Name] -> This will return the data assigned to the key
	fmt.Println(mpp["Sahil"])

	//How to add new value to the map
	// map_name[New_Key_Name_You_Want_To_Add] = value_OR_data_you_want_to_store
	mpp["New_Data"] = 1001
	fmt.Println(mpp["New_Data"])

	/*
		Comma Ok Idiom
		-> It is used to check whether the key is present in the map or not

		Syntax of the this is below:
			* value, ok := map_name[Key_Name]
			//  For value --> If the value is present in the map then it will return the value and if the value is not presemt in the map then it will return 0
			//  For Ok --> If the key is present in the map then i will return the true and if the key is not present in the map then  it will return false.
	*/

	value, ok := mpp["Sahil"]

	// Value will be the value assigned to the key Sahil
	// Ok will be true as Sahil is present in the map

	fmt.Println(value, ok)

	value, ok = mpp["Sahil1"]

	// Value will be the 0 as this value is not present in the map
	// Ok will be the false as this key is not present in the map

	fmt.Println(value, ok)

	/*
		Deleting th keys from the map
		-> we can use delete() function to delete the key from the map
		-> Syntax of the delete() function is below:
		delete(map_name, key_name)
		-> This function will take two arguments:
		-> First argument is the map name from where we want to delete the key
		-> Second argument is the key name that we want to delete from the map
	*/

	delete(mpp, "Sahil")
	value, ok = mpp["Sahil"]
	// Value will be the 0 as we delete this value and ok will be false.
	fmt.Println(value, ok) // Output : 0 false

	// We Can also empty the map with the help of the another built-in function clear(map_name)
	// It will take the map_name  as the argument which we will want to empty
	clear(mpp)

	fmt.Println(mpp) // Output : mpp[]
	// The Map will become empty

	// We can also check if the two maps are equal or not with the help of one other built-in function
	/*
		maps.equal(map_name_1, map_name_2)
		-> This function will return the  boolean valu
		-> This will take the two maps as the arguments and check theirs keys and values to each other
		-> If every value and key of these two maps are equal and identical then it will return the true otherwise it will return false
	*/

	mpp1 := map[string]int{}

	fmt.Println(maps.Equal(mpp, mpp1))
	// Output : true
	// As both the maps are empty at a moment

}
