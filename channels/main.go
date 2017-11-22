package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link) // The `_` here is because our program does not need the `resp` object, only the `err` object

	if err != nil {
		fmt.Println(link, "might be down!")
		return // This will exit out of the if statement
	}

	fmt.Println(link, "is up!")
}
