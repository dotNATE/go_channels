package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://go.dev",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for {
		go checkUrl(<-c, c)
	}
}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s might be down! \n", url)
		c <- url
		return
	}
	fmt.Printf("%s seems to be okay! \n", url)
	c <- url
}
