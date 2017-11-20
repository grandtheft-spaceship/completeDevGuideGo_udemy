package main

import "fmt"

type englishBot struct{} // We do need any properties // We're making them structs just so we can add some functionality
type spanishBot struct{}

func (englishBot) getGreeting() string { // Because we're not creating full functionality, we can leave off the "eb" we'd normally put to describe the receiver and simply put the *type* of receiver we expect so we do not see the error message from the compiler of defining a variable and not using the variable
	// Custom logic to generate an English greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// Custom logic to generate a Spanish greeting
	return "¡Holla!"
}

func printGreeting(eb englishBot) { // Our goal here is to use a single `printGreeting()` function for both `eb` and `sb`
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) { // When first saving the file, we get the error that says we have defined the same function twice, even if they take two different arguments
	fmt.Println(sb.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
