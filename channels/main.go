package main

import (
	"fmt"
	"net/http"
	"time"
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

	// for { // Infinite loop syntax
	// 	go checkLink(<-c, c) // `go` keyword is used here to continue pinging links until we get an error // <- c is a link and c represents a channel
	// }

	// Another syntax format as the for loop above
	for l := range c { // Wait for the channel to return some value; when it has, place the value into the variable `l`
		go func(link string) { // Function Literal // passing in `link` will update the value of `link`
			time.Sleep(5 * time.Second) // Pauses current goroutine for 5 seconds
			// checkLink(l, c) // We get a warning here because of what we learned about pointers and RAM memory / pass-by value in Go
			checkLink(link, c)
		}(l) // This () at the end invokes the function // passing in `l` will update the value of `link`
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // The `_` here is because our program does not need the `resp` object, only the `err` object

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // Feeding each link into the channel
		return    // This will exit out of the if statement
	}
	fmt.Println(link, "is up!")
	c <- link // Feeding each link into the channel
}
