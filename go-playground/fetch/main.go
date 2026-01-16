package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url  string
	size int64
	err  error
}

func fetch(url string, out chan<- result) {
	res, err := http.Get(url)
	if err != nil {
		out <- result{url: url, size: 0, err: err}
		return
	}
	defer res.Body.Close()

	out <- result{
		url:  url,
		size: res.ContentLength,
		err:  nil,
	}

}
func main() {
	var result = make(chan result, 4)
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.golang.org",
		"https://www.wikipedia.org",
	}

	for _, url := range urls {
		go fetch(url, result)
	}
	
	winner := <-result

	if winner.err != nil {
		fmt.Printf("The first responder (%s) failed with error: %v\n", winner.url, winner.err)
	} else {
		fmt.Printf("Winner: %s with size %d bytes\n", winner.url, winner.size)
	}

	fmt.Println("Main is exiting. Remaining goroutines will finish in the background.")

}
