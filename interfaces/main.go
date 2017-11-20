package main

import "fmt"

type bot interface { //
	getGreeting() string
}

type englishBot struct{} // We do need any properties // We're making them structs just so we can add some functionality
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string { // Because we're not creating full functionality, we can leave off the "eb" we'd normally put to describe the receiver and simply put the *type* of receiver we expect so we do not see the error message from the compiler of defining a variable and not using the variable
	// Custom logic to generate an English greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// Custom logic to generate a Spanish greeting
	return "¡Holla!"
}
