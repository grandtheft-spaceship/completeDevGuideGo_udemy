package main

import "fmt"

func main() {
	// var colors map[string]string // Another way to declare a map

	// colors := make(map[string]string) // Another way to declare a map // make() is a built-in function in Go
	// colors["white"] = "#fffff"
	// delete(colors, "white") // delete() is a built-in function

	colors := map[string]string{ //Creating a map where all the keys are type string and all the values are type string
		"red":   "#ff0000",
		"green": "#4bf745", //Unlike other languages, you need a comma even on the last key-value pair
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
