package main

import "fmt"

type person struct { // Creates a new type, "person", that is really of type "struct"
	firstName string // First name property for type "person"
	lastName  string // Last name property for type "person"
}

func main() {
	// alex := person{"Alex", "Anderson"} // Relies on order of definition of fields for assignment // NOT RECOMMENDED WAY
	// alex := person{firstName: "Alex", lastName: "Anderson"} // Another way for defining a struct // MORE RELIABLE
	var alex person            // Third way to declare a new struct // Properties will be of "zero value"
	alex.firstName = "Alex"    // Updates firstName property
	alex.lastName = "Anderson" // Updates lastName property
	// fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
