package main

import "fmt"

/*
	Embedded Struct
	When we want to use the same struct in other structs, we can embedded one struct into another.

	See The Below Example:
*/

type Address struct {
	street     string
	city       string
	country    string
	postalCode int
}

type Pesrons struct {
	firstName string
	lastName  string
	age       int
	Address   // Now The Struct Person also has all the fields of the Address struct.
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
		Address: Address{
			street:     "123 Garhshankar",
			city:       "Jalandhar",
			country:    "India",
			postalCode: 144001,
		},
	}
	sahil.personInfo()
	sahil.updateName("King", "Thakur")
	// Output : FirstName: King, LastName: Thakur

	sahil.updateNameCopy(24)
	fmt.Println(sahil.age)
	// Output : 23

	fmt.Println(sahil.Address.city)
	// Output : Jalandhar

}

/*
	Pointer Receiver vs Value Receiver

	Pointer Receiver methods are those methods which are attached to the pointer of the struct.
	They can change the original value of the struct or the receiver

	func(p *Person) updateName(newName string){
		p.name = name
		fmt.Printf("Updated Name: %s", p.name)

		This way we can change the original value of the receiver. This is called Pointer Receiver.ss
	}
*/

func (p *Pesrons) updateName(newFirstName, newLastName string) {
	p.firstName = newFirstName
	p.lastName = newLastName
	fmt.Println(p.firstName + " " + p.lastName)
}

/*
	Value Receiver methods take the parameter as a copy of the original value.
	So, if we even change the value of the receiver, it will still not change the original value.

	func (p Person) updateNameCopy(age int){
		p.age = age
		fmt.Printf("Updated Age: %d", p.age)
	}
*/

func (p Pesrons) updateNameCopy(age int) {
	p.age = age
}
