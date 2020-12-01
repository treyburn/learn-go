package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://totallyfake.abc",
		"http://alteryx.com",
		"http://reddit.com",
		"http://colorado.edu",
		"http://www.solveigdelabroye.me",
	}

	c := make(chan string)

	for _, u := range urls {
		go checkUrl(u, c)
	}

	for i := 0; i < len(urls); i ++ {
		fmt.Println(<-c)
	}
}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		c <- "Status: "+url+" might be down."
		return
	}

	c <- "Status: "+url+" is up."
}