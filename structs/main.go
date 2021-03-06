package main

import "fmt"

type person struct { // Creates a new type, "person", that is really of type "struct"
	firstName string // First name property for type "person"
	lastName  string // Last name property for type "person"
	// contact   contactInfo
	contactInfo // Another way to embed a struct property - This says that we have a field name of contactInfo of type contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{"Alex", "Anderson"} // Relies on order of definition of fields for assignment // NOT RECOMMENDED WAY

	// alex := person{firstName: "Alex", lastName: "Anderson"} // Another way for defining a struct // MORE RELIABLE

	// var alex person            // Third way to declare a new struct // Properties will be of "zero value"
	// alex.firstName = "Alex"    // Updates firstName property
	// alex.lastName = "Anderson" // Updates lastName property
	// alex.contact.email = "aanderson@gmail.com"
	// alex.contact.zipCode = 12345

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		// contact: contactInfo{
		contactInfo: contactInfo{
			email:   "aanderson@gmail.com",
			zipCode: 12345,
		},
	}
	// fmt.Println(alex)
	// alex.print()
	// alexPointer := &alex                // Finds the memory address of alex and stores it in alexPointer

	alex.updateName("Alejandro") // The updateName() function will still work because Go automatically converts alex into alexPointer for us in the updateName function
	alex.print()                 // Run this code and you'll see that firstName property did not update
}

func (pointerToPerson *person) updateName(newFirstName string) { // *person is a type description - it means we're working with a pointer to a person
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
