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

	c := make(chan string) //create channel with type of string

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		// this part will block the main routine until receive message from channel
		go func(link string) { // this is same like `use` in PHP when we want to pass variable to anonymous function block
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // this is function literal in other language it called anonymous function
	}
}

func checkLink(link string, c chan string) {
	// blocking operation
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")

	c <- link
}
