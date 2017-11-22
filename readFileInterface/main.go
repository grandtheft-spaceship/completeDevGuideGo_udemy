package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
	// fn := os.Args[1] // IMPLEMENTATION v1
	// fmt.Println(readFile(fn)) // IMPLEMENTATION v1
}

// func readFile(filename string) string { // IMPLEMENTATION v1
// 	s, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}
// 	return string(s)
// }
