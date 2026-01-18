package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type Result struct {
	url   string
	title string
	err   error
}

func getTitle(url string) (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return "", fmt.Errorf("title not found")
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "title" {
				tokenizer.Next()
				return tokenizer.Token().Data, nil
			}
		}
	}
}

func main() {
	urls := []string{
		"https://google.com",
		"https://bhanuprakashsai.com",
		"https://pkg.go.dev/net/http",
		"https://go.dev",
		"https://o.dev",
	}
	var wg sync.WaitGroup
	var resultsChan = make(chan Result, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			title, err := getTitle(url)

			resultsChan <- Result{
				url:   url,
				title: title,
				err:   err,
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for res := range resultsChan {
		if res.err != nil {
			fmt.Printf("[Error] %s: %v\n", res.url, res.err)
		} else {
			fmt.Printf("[Success] %s -> %s\n", res.url, res.title)
		}
	}

	fmt.Println("[Done]")

}
