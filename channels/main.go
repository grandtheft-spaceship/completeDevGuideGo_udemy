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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // The `_` here is because our program does not need the `resp` object, only the `err` object

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think" // Inputting data into the channel
		return                       // This will exit out of the if statement
	}
	fmt.Println(link, "is up!")
	c <- "Yup, it's up" // Inputting data into the channel
}
