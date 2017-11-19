package main

import "fmt"

func main() {
	colors := map[string]string{ //Creating a map where all the keys are type string and all the values are type string
		"red":   "#ff0000",
		"green": "#4bf745", //Unlike other languages, you need a comma even on the last key-value pair
	}

	fmt.Println(colors)
}
