package main

import (
	"encoding/json"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	mpp := make([]Person, 0)

	mpp = append(mpp, Person{Name: "John", Age: 30})
	mpp = append(mpp, Person{Name: "Jane", Age: 25})
	mpp = append(mpp, Person{Name: "Doe", Age: 40})

	val, err := json.Marshal(mpp)
	if err != nil {
		panic(err)
	}
	println(string(val))
}
