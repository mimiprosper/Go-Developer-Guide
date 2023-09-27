package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// slice of strings
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://facebook.com",
		"http://golang.com",
	}

	// create a channel
	c := make(chan string)

	for _, link := range links {
		// create 'goroutine' with a channel as an argument (concurency)
		go checkLink(link, c)
	}

	// repeating goroutine
	// for {
	// 	go checkLink(<-c, c)
	// }

	// repeating goroutine & pause every 5s
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)

	}

}

// check link status and print site status with a channel as a argument
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, " is up")
	c <- link
}
