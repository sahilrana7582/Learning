package main

import "fmt"

type Pesrons struct {
	firstName string
	lastName  string
	age       int
}

/*
	This is Not Function.
	This is Methos
	Method are almost same as function but they are attached to the struct.
	Method are different in syntax as they receives a receiver argument between the func keyword and the function name.
	func (receiver_name receiver_type) method_name(parameters) (return_type){}

*/

func (p Pesrons) personInfo() {
	fmt.Println("First Name: %s, Last Name: %s, Age: %s", p.firstName, p.lastName, p.age)
}

func main() {
	sahil := Pesrons{
		firstName: "Sahil",
		lastName:  "Rana",
		age:       23,
	}
	sahil.personInfo()

}
